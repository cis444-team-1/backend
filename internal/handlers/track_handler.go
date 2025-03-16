package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (*Handler) InsertTrackHandler(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Track was successfully created",
		"trackId": "ID NUMBER",
	})
}

func (*Handler) UpdateTrackHandler(c echo.Context) error {
	return nil
}

func (*Handler) DeleteTrackHandler(c echo.Context) error {
	return nil
}
