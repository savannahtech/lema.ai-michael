package server

import (
	"context"
	"encoding/json"
	"github.com/dilly3/houdini/internal/repository"
	"github.com/dilly3/houdini/internal/repository/cache"
	"github.com/dilly3/houdini/internal/server/response"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strconv"
)

type SettingsPayload struct {
	Owner   string `json:"owner" validate:"required"`
	Repo    string `json:"repo" validate:"required"`
	Since   string `json:"since" validate:"required"`
	PerPage int    `json:"per_page" validate:"required"`
}

func (h *Handler) UpdateSettingsHandler(w http.ResponseWriter, r *http.Request) {
	// Read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Err(err).Msg("failed to close body")
			return
		}
	}(r.Body)

	// Unmarshal the JSON data into the Payload struct
	var payload SettingsPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cac := cache.GetDefaultCache()
	cac.SetOwner(payload.Owner)

	cac.SetRepo(payload.Repo)
	cac.SetSince(payload.Since)
	inter := strconv.Itoa(payload.PerPage)
	cac.SetPerPage(inter)
	// Delete all commits newer than the new since date
	ctx := context.Background()
	err = repository.GetDefaultStore().DeleteByDate(ctx, cac.GetRepo(), cac.GetSince())
	if err != nil {
		h.Logger.Error().Err(err).Msg("failed to delete commits")
	}
	response.RespondWithJson(w, "settings updated successfully", http.StatusOK, nil)
}
