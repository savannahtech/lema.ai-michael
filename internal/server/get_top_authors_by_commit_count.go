package server

import (
	"context"
	"github.com/dilly3/houdini/internal/repository"
	"github.com/dilly3/houdini/internal/server/response"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) GetTopAuthorsByCommitsHandler(w http.ResponseWriter, r *http.Request) {
	repoName := chi.URLParam(r, "repo_name")
	limitString := chi.URLParam(r, "limit")
	if limitString == "" {
		limitString = "10"
	}
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		h.Logger.Error().Err(err).Msg("failed to convert limit to int")
		response.RespondWithError(w, http.StatusBadRequest, "limit must be a number")
		return
	}

	if repoName == "" {
		h.Logger.Error().Msg("repo name is required")
		response.RespondWithError(w, http.StatusBadRequest, "repo name is required")
		return
	}
	ctx := context.Background()
	res, err := repository.GetDefaultStore().GetTopCommitsAuthorsByCount(ctx, repoName, limit)
	if err != nil {
		h.Logger.Error().Err(err).Msg("failed to get top authors by commits")
		response.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	if len(res) == 0 {
		response.RespondWithJson(w, "no authors found", http.StatusOK, res)
		return
	}
	var message string
	if len(res) == 0 {
		message = "no authors found"
	} else {
		message = "top authors by commits retrieved"
	}
	response.RespondWithJson(w, message, http.StatusOK, res)

}
