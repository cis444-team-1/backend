package routes

import (
	"net/http"

	"github.com/cis444-team-1/backend/internal/auth"
	"github.com/cis444-team-1/backend/internal/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, h *handlers.Handler) {
	e.POST("/generate-presigned-url", h.GeneratePresignedFileURL, auth.AuthMiddleware)
	e.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	}, auth.AuthMiddleware)
}
