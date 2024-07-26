package main

import (
	"github.com/dilly3/houdini/internal/config"
	ghi "github.com/dilly3/houdini/internal/github"
	"github.com/dilly3/houdini/internal/repository"
	"github.com/dilly3/houdini/internal/repository/cache"
	"github.com/dilly3/houdini/internal/server"
	"github.com/dilly3/houdini/pkg/cron"
	"github.com/dilly3/houdini/pkg/github"
	"github.com/dilly3/houdini/storage/postgres"
	"github.com/dilly3/houdini/storage/redis"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	err := config.LoadConfig(".env", &logger)
	if err != nil {
		logger.Error().Err(err).Msgf("Error loading .env file: %v", err)
		os.Exit(1)
	}
	c := config.Config
	pgs := postgres.New(c, &logger)
	repository.NewStore(pgs)
	settings := config.GetSettings()
	redisClt, err := redis.NewRedisClient(c.RedisADDR, c.RedisUser, c.RedisPassword, settings, &logger)
	if err != nil {
		logger.Error().Err(err).Msgf("redis connection: %s", err.Error())
		os.Exit(1)
	}
	cache.NewCache(redisClt)
	githubClient := github.NewGHClient(c.GithubBaseURL, c.GithubToken)
	gitHubInteract := ghi.NewGHubITR(githubClient)
	handler := server.NewHandler(&logger)
	cron.InitCron()
	cron.SetCronJob(gitHubInteract.GetCommitsCron, cron.GetTimeDuration(c.CronInterval))
	cron.SetCronJob(gitHubInteract.GetRepoCron, cron.GetTimeDuration(c.CronInterval))
	go cron.StartCronJob()
	httpHandler := server.NewChiRouter(handler, time.Minute)
	httpServer := &http.Server{
		Addr:    config.Config.Port,
		Handler: httpHandler,
	}
	go server.GetLimiter().CleanUp()
	logger.Info().Msgf("Server started on port %s", c.Port)
	if err := httpServer.ListenAndServe(); err != nil {
		cron.StopCronJob()
		logger.Error().Err(err).Msg(err.Error())
	}
}
