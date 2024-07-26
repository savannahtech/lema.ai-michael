package server

import (
	"github.com/dilly3/houdini/internal/github"
	"github.com/dilly3/houdini/internal/server/response"
	"net/http"
	"strconv"
)

func (h *Handler) ListCommitsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	repo := params.Get("repo")
	if repo == "" {
		response.RespondWithError(w, http.StatusBadRequest, "repo is required")
		return
	}
	owner := params.Get("owner")
	if owner == "" {
		response.RespondWithError(w, http.StatusBadRequest, "owner is required")
		return
	}
	since := params.Get("since")
	page := params.Get("page")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, "page is required as integer")
		return
	}
	getCommits, err := github.GetGitHubAdp().ListCommits(owner, repo, since, intPage)
	if err != nil {
		h.Logger.Error().Err(err).Msg("failed to list commits")
		response.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.RespondWithJson(w, "commits retrieved successfully", http.StatusOK, getCommits)
	return
}
