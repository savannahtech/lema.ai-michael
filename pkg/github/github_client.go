package github

import (
	"net/http"
	"time"
)

type GHClient struct {
	BaseURL    string
	token      string
	HTTPClient *http.Client
}

func NewGHClient(baseUrl, token string) *GHClient {
	return &GHClient{
		BaseURL: baseUrl,
		token:   token,
		HTTPClient: &http.Client{
			Timeout: 1 * time.Minute,
		},
	}
}
