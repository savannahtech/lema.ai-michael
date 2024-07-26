package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var (
	owner   = "github_owner"
	repo    = "github_repo"
	since   = "github_since"
	perPage = "github_per_page"
)

type Configuration struct {
	Port             string
	Env              string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresTimezone string
	GithubBaseURL    string
	GithubPerPage    string
	GithubSince      string
	GithubToken      string
	GithubOwner      string
	GithubRepo       string
	CronInterval     string
	NetworkRetry     int
	RedisHost        string
	RedisADDR        string
	RedisPassword    string
	RedisUser        string
}

var Config *Configuration

func LoadConfig(envFile string, logger *zerolog.Logger) error {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	if envFile == "" {
		envFile = ".env"
	}
	log.Printf("sourcing %v", envFile)
	err := godotenv.Load(fmt.Sprintf("%s/../../%s", basePath, envFile))
	if err != nil {
		logger.Error().Err(err).Msgf("Error loading .env file: %v", err)
		return err
	}
	Config = &Configuration{}
	Config.Port = os.Getenv("PORT")

	Config.PostgresHost = os.Getenv("POSTGRES_HOST")
	Config.PostgresPort = os.Getenv("POSTGRES_PORT")
	Config.PostgresUser = os.Getenv("POSTGRES_USER")
	Config.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	Config.PostgresDB = os.Getenv("POSTGRES_DB")
	Config.PostgresTimezone = os.Getenv("POSTGRES_TIMEZONE")
	Config.GithubBaseURL = os.Getenv("GITHUB_BASE_URL")
	Config.GithubPerPage = os.Getenv("GITHUB_PER_PAGE")
	Config.GithubSince = os.Getenv("GITHUB_SINCE")
	Config.GithubToken = os.Getenv("GITHUB_TOKEN")
	Config.GithubOwner = os.Getenv("GITHUB_OWNER")
	Config.GithubRepo = os.Getenv("GITHUB_REPO")
	Config.CronInterval = os.Getenv("CRON_INTERVAL")
	Config.NetworkRetry, err = strconv.Atoi(os.Getenv("NETWORK_RETRY"))
	if err != nil {
		logger.Error().Err(err).Msgf("Error converting network retry to int: %v", err)
		return err
	}
	Config.RedisPassword = os.Getenv("REDIS_PASSWORD")
	Config.RedisUser = os.Getenv("REDIS_USER")
	Config.RedisADDR = os.Getenv("REDIS_ADDR")
	return nil

}

func GetSettings() map[string]string {
	return map[string]string{
		owner:   Config.GithubOwner,
		repo:    Config.GithubRepo,
		since:   Config.GithubSince,
		perPage: Config.GithubPerPage,
	}
}

func GetTimeDuration() time.Duration {
	aInt, err := strconv.Atoi(Config.CronInterval)
	if err != nil {
		log.Println("failed to convert cron interval to int")
		return 1
	}
	return time.Minute * time.Duration(aInt)

}
