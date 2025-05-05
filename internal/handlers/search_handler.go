package handlers

import (
	"net/http"

	"github.com/cis444-team-1/backend/internal/db/repositories"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SearchHandler(c echo.Context) error {
	query := c.QueryParam("q")

	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot parse query"})
	}

	dbConn := h.DB.GetDB()
	trackRepo := repositories.NewTrackRepository(dbConn)
	playlistRepo := repositories.NewPlaylistRepository(dbConn)

	tracks, err := trackRepo.GetTracksBySearchQuery(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"errors": err.Error()})
	}

	playlists, err := playlistRepo.GetPlaylistsBySearchQuery(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"errors": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"tracks":    tracks,
		"playlists": playlists,
	})
}
