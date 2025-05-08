package handlers

import (
	"net/http"

	"github.com/cis444-team-1/backend/internal/db/repositories"
	"github.com/cis444-team-1/backend/internal/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUserHandler(c echo.Context) error {
	userId := c.Param("userId")

	if userId == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot parse user id"})
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user id"})
	}

	dbConn := h.DB.GetDB()
	userRepo := repositories.NewUserRepository(dbConn)
	user, err := userRepo.GetUserByUserID(userUUID.String())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":            user.ID,
		"user_metadata": user.UserMetadata,
	})
}

func (h *Handler) GetUsersByUserIDsHandler(c echo.Context) error {

	var userIds models.UserIds
	err := c.Bind(&userIds)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	dbConn := h.DB.GetDB()
	userRepo := repositories.NewUserRepository(dbConn)
	users, err := userRepo.GetUsersByUserIDs(userIds)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}
