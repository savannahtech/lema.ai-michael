package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"net/http"
	"time"
)

func NewChiRouter(h *Handler, limiterDuration time.Duration) *chi.Mux {
	router := chi.NewRouter()
	limiter := NewRateLimiter(limiterDuration)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allows all origins
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})

	router.Use(c.Handler)
	// Add Logger to router
	routerWithLog := router.With(h.loggingMiddleware)
	routerWithLog.Patch("/v1/settings", h.UpdateSettingsHandler)

	// Add rate limit middleware
	limitRoutes := routerWithLog.With(limiter.IPRateLimit)

	limitRoutes.Get("/v1/repo", h.GetRepoHandler)
	limitRoutes.Get("/v1/commits", h.ListCommitsHandler)
	limitRoutes.Get("/v1/repo/{name}", h.GetRepoByName)
	limitRoutes.Get("/v1/commits/{name}/{limit}", h.GetCommitsByRepoName)
	limitRoutes.Get("/v1/repos/{language}/{limit}", h.GetReposByLanguage)
	limitRoutes.Get("/v1/repos-stars/{limit}", h.GetRepoByStarsCount)
	limitRoutes.Get("/v1/authors/top/{repo_name}/{limit}", h.GetTopAuthorsByCommitsHandler)

	return router
}
