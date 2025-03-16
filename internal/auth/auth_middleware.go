package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/cis444-team-1/backend/config"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cfg := config.LoadConfig()

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		// Validate the user's token
		token := strings.Replace(authHeader, "Bearer ", "", 1)

		user, err := cfg.SupabaseClient.Auth.User(context.Background(), token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		c.Set("user", user)
		return next(c)
	}
}
