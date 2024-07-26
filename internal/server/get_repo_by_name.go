package server

import (
	"github.com/dilly3/houdini/internal/repository"
	"github.com/dilly3/houdini/internal/server/response"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) GetRepoByName(w http.ResponseWriter, r *http.Request) {
	repoName := chi.URLParam(r, "name")
	if repoName == "" {
		h.Logger.Error().Msg("repo name is required")
		response.RespondWithError(w, http.StatusBadRequest, "repo name is required")
		return
	}

	repo, err := repository.GetDefaultStore().GetRepoByName(r.Context(), repoName)
	if err != nil {
		h.Logger.Error().Err(err).Msg(err.Error())
		response.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	message := "repo retrieved successfully"
	if repo == nil {
		message = "no repo found"
	}

	response.RespondWithJson(w, message, http.StatusOK, repo)
	return
}
