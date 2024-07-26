package server

import (
	"github.com/dilly3/houdini/internal/repository"
	"github.com/dilly3/houdini/internal/server/response"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) GetReposByLanguage(w http.ResponseWriter, r *http.Request) {
	language := chi.URLParam(r, "language")
	if language == "" {
		h.Logger.Error().Msg("language is required")
		response.RespondWithError(w, http.StatusBadRequest, "language is required")
		return
	}
	limit := chi.URLParam(r, "limit")
	if limit == "" {
		h.Logger.Error().Msg("limit is required")
		response.RespondWithError(w, http.StatusBadRequest, "limit is required")
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		h.Logger.Error().Err(err).Msg("failed to convert limit to int")
		response.RespondWithError(w, http.StatusBadRequest, "limit should be an integer")
		return
	}

	repos, err := repository.GetDefaultStore().GetReposByLanguage(r.Context(), language, limitInt)
	if err != nil {
		h.Logger.Error().Err(err).Msg("failed to get repos by language")
		response.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	message := "repos retrieved successfully"
	if len(repos) == 0 {
		message = "no repos found"
	}

	response.RespondWithJson(w, message, http.StatusOK, repos)
	return
}
