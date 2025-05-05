package handlers

import (
	"net/http"

	"github.com/cis444-team-1/backend/internal/db/repositories"
	"github.com/cis444-team-1/backend/internal/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (h *Handler) GetPlaylistHandler(c echo.Context) error {
	playlistId := c.Param("playlistId")

	if playlistId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot parse playlist id"})
	}

	playlistUUID, err := uuid.Parse(playlistId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid playlist id"})
	}

	dbConn := h.DB.GetDB()
	playlistRepo := repositories.NewPlaylistRepository(dbConn)
	playlist, err := playlistRepo.GetPlaylistByID(playlistUUID.String())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, playlist)
}

func (h *Handler) DeletePlaylistHandler(c echo.Context) error {
	playlistId := c.Param("playlistId")

	if playlistId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot parse playlist id"})
	}

	playlistUUID, err := uuid.Parse(playlistId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid playlist id"})
	}

	dbConn := h.DB.GetDB()
	playlistRepo := repositories.NewPlaylistRepository(dbConn)

	err = playlistRepo.DeletePlaylist(playlistUUID.String())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message":    "Playlist was successfully deleted",
		"playlistId": playlistId,
	})
}

func (*Handler) UpdatePlaylistHandler(c echo.Context) error {
	return nil
}

func (h *Handler) InsertPlaylistHandler(c echo.Context) error {
	// user := c.Get("user")

	// if (user == nil) {
	// 	return c.JSON(http.StatusUnauthorized, "Unauthorized")
	// }

	var playlist models.PostPlaylistRequest

	// Bind JSON request to playlist struct
	if err := c.Bind(&playlist); err != nil {
		return err
	}

	dbConn := h.DB.GetDB()
	playlistRepo := repositories.NewPlaylistRepository(dbConn)

	playlistID, err := playlistRepo.CreatePlaylist(uuid.New().String(), playlist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"errors": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Playlist was successfully created", "playlist_id": playlistID})
}

func (h *Handler) AddTrackToPlaylistHandler(c echo.Context) error {
	playlistId := c.Param("playlistId")

	if playlistId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot parse playlist id"})
	}

	playlistUUID, err := uuid.Parse(playlistId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid playlist id"})
	}

	var track models.AddTrackToPlaylistRequest

	// Bind JSON request to playlist struct
	if err := c.Bind(&track); err != nil {
		return err
	}

	dbConn := h.DB.GetDB()
	playlistRepo := repositories.NewPlaylistRepository(dbConn)

	err = playlistRepo.AddTrackToPlaylist(playlistUUID.String(), track.TrackId.String(), track.Position)

	if err != nil {
		if err.(*pq.Error).Code == "23505" {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Track already exists in playlist"})
		}

		if err.(*pq.Error).Code == "23503" {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Track does not exist"})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{"errors": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Track was successfully added to playlist", "playlist_id": playlistId, "track_id": track.TrackId.String()})
}

func (h *Handler) RemoveTrackFromPlaylistHandler(c echo.Context) error {
	playlistId := c.Param("playlistId")

	if playlistId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot parse playlist id"})
	}

	playlistUUID, err := uuid.Parse(playlistId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid playlist id"})
	}

	var track models.RemoveTrackFromPlaylistRequest

	// Bind JSON request to playlist struct
	if err := c.Bind(&track); err != nil {
		return err
	}

	dbConn := h.DB.GetDB()
	playlistRepo := repositories.NewPlaylistRepository(dbConn)

	err = playlistRepo.RemoveTrackFromPlaylist(playlistUUID.String(), track.TrackId.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"errors": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Track was successfully removed from playlist", "playlist_id": playlistId, "track_id": track.TrackId.String()})
}

func (h *Handler) GetTracksFromPlaylistHandler(c echo.Context) error {
	playlistId := c.Param("playlistId")

	if playlistId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot parse playlist id"})
	}

	playlistUUID, err := uuid.Parse(playlistId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid playlist id"})
	}

	dbConn := h.DB.GetDB()
	playlistRepo := repositories.NewPlaylistRepository(dbConn)

	tracks, err := playlistRepo.GetTracksByPlaylistID(playlistUUID.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"errors": err.Error()})
	}

	return c.JSON(http.StatusOK, tracks)
}
