package server

import "github.com/rs/zerolog"

type Handler struct {
	Logger *zerolog.Logger
}

func NewHandler(logger *zerolog.Logger) *Handler {
	return &Handler{
		Logger: logger,
	}
}
