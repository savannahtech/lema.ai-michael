package server

import (
	"encoding/json"
	"github.com/dilly3/houdini/internal/config"
	gh "github.com/dilly3/houdini/internal/github"
	"github.com/dilly3/houdini/internal/server/response"
	"github.com/dilly3/houdini/pkg/github"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var handler *Handler

func TestMain(m *testing.M) {
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	err := config.LoadConfig(".env", &logger)
	if err != nil {
		logger.Error().Err(err).Msgf("Error loading .env file: %v", err)
		os.Exit(1)
	}
	c := config.Config
	ghc := github.NewGHClient(c.GithubBaseURL, c.GithubToken)
	gh.NewGHubITR(ghc)
	handler = NewHandler(&logger)
	exitCode := m.Run()

	os.Exit(exitCode)
}

func executeGetRequest(pattern, path string, handlr http.HandlerFunc) (*response.HttpResponse, int, error) {

	r := chi.NewRouter()
	r.Get(pattern, handlr)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, 0, err
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Serve the HTTP request using the router
	r.ServeHTTP(rr, req)
	res := &response.HttpResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), res)
	if err != nil {
		return nil, 0, err
	}
	return res, rr.Code, nil
}
