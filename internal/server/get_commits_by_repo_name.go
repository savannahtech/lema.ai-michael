package server

import (
	"github.com/dilly3/houdini/internal/repository"
	"github.com/dilly3/houdini/internal/server/response"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) GetCommitsByRepoName(w http.ResponseWriter, r *http.Request) {
	repoName := chi.URLParam(r, "name")
	limit := chi.URLParam(r, "limit")
	if repoName == "" {
		h.Logger.Error().Msg("repo name is required")
		response.RespondWithError(w, http.StatusBadRequest, "repo name is required")
		return
	}
	if limit == "" {
		h.Logger.Error().Msg("limit is required")
		response.RespondWithError(w, http.StatusBadRequest, "limit is required")
		return
	}
	// convert limit to int
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		h.Logger.Error().Err(err).Msg("failed to convert limit to int")
		response.RespondWithError(w, http.StatusBadRequest, "limit should be an integer")
		return
	}

	commits, err := repository.GetDefaultStore().GetCommitsByRepoName(r.Context(), repoName, limitInt)
	if err != nil {
		h.Logger.Error().Err(err).Msg("failed to get commits by repo name")
		response.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	message := "commits retrieved successfully"
	if len(commits) == 0 {
		message = "no commits found"
	}

	response.RespondWithJson(w, message, http.StatusOK, commits)
	return
}
