package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cis444-team-1/backend/config"
	"github.com/cis444-team-1/backend/internal/db"
	"github.com/cis444-team-1/backend/internal/handlers"
	"github.com/cis444-team-1/backend/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dbConn := db.GetInstance()
	defer dbConn.Close()

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] method=${method}, uri=${uri}, status=${status}, latency=${latency_human}, error=${error}\n",
	}))
	e.Use(middleware.Recover())

	config.InitCORS(e)

	// Intialize API routes
	h := handlers.NewHandler(dbConn)
	routes.InitRoutes(e, h)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			log.Fatalf("Error shutting down server: %v", err)
		}
	}()

	log.Println("Starting server on port 8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
