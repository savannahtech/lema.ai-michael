package server

import (
	"net/http"
	"time"
)

func (h *Handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.Logger.Info().Msgf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		h.Logger.Info().Msgf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
