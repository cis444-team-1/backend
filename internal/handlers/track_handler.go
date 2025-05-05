package handlers

import (
	"net/http"

	"github.com/cis444-team-1/backend/internal/db/repositories"
	"github.com/cis444-team-1/backend/internal/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetTrackByTrackIDHandler(c echo.Context) error {
	trackId := c.Param("trackId")

	if trackId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot parse track id"})
	}

	trackUUID, err := uuid.Parse(trackId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid track id"})
	}

	dbConn := h.DB.GetDB()
	trackRepo := repositories.NewTrackRepository(dbConn)
	track, err := trackRepo.GetTrackByID(trackUUID.String())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, track)
}

func (h *Handler) InsertTrackHandler(c echo.Context) error {
	// user := c.Get("user")

	// if (user == nil) {
	// 	return c.JSON(http.StatusUnauthorized, "Unauthorized")
	// }

	var track models.PostTrackRequest

	// Bind JSON request to track struct
	if err := c.Bind(&track); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	dbConn := h.DB.GetDB()
	trackRepo := repositories.NewTrackRepository(dbConn)

	trackID, err := trackRepo.CreateTrack(uuid.New().String(), track)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Track was successfully created", "track_id": trackID})
}

func (*Handler) UpdateTrackHandler(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Track was successfully created",
		"trackId": "ID NUMBER",
	})
}

func (h *Handler) DeleteTrackHandler(c echo.Context) error {
	trackId := c.Param("trackId")

	if trackId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot parse track id"})
	}

	trackUUID, err := uuid.Parse(trackId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid track id"})
	}

	dbConn := h.DB.GetDB()
	trackRepo := repositories.NewTrackRepository(dbConn)

	err = trackRepo.DeleteTrack(trackUUID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Track was successfully created",
		"trackId": "ID NUMBER",
	})
}
