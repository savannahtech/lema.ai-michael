package server

import (
	"github.com/dilly3/houdini/internal/repository"
	"github.com/dilly3/houdini/internal/server/response"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) GetRepoByStarsCount(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(chi.URLParam(r, "limit"))
	if err != nil {
		h.Logger.Error().Err(err).Msg("failed to convert limit to int")
		response.RespondWithError(w, http.StatusBadRequest, "limit must be a number")
		return
	}

	repos, err := repository.GetDefaultStore().GetReposByStarCount(r.Context(), limit)
	if err != nil {
		h.Logger.Error().Err(err).Msg("GetRepoByStarCount:failed to get repo")
		response.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	message := "repo retrieved successfully"
	if len(repos) == 0 {
		message = "no repo found"
	}

	response.RespondWithJson(w, message, http.StatusOK, repos)
	return
}
