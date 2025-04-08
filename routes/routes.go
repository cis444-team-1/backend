package routes

import (
	"github.com/cis444-team-1/backend/internal/auth"
	"github.com/cis444-team-1/backend/internal/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, h *handlers.Handler) {
	// -- File Upload Handlers
	e.POST("/generate-presigned-url", h.GeneratePresignedFileURL, auth.AuthMiddleware)

	// -- Track Handlers
	e.GET("/tracks/uploads", nil)                      // Get user's uploads, Has query params for limit and offset
	e.POST("/tracks", h.InsertTrackHandler)            // Create a new track, Requires middleware
	e.GET("/tracks/:trackId", nil)                     // Get specific track by its id
	e.PUT("/tracks/:trackId", h.UpdateTrackHandler)    // Update a track by its id
	e.DELETE("/tracks/:trackId", h.DeleteTrackHandler) // Delete a track by its id
	e.GET("/tracks/top-charts", nil)                   // Get All-time popular songs
	e.GET("/tracks/trending", nil)                     // Get Recently popular songs
	e.GET("/tracks/genres/:genre", nil)                // Get tracks by genre with query params
	e.GET("/tracks/most-played", nil)                  // Get user's most played songs, Requires middleware
	e.GET("/tracks/history", nil)                      // Get user's play history, Requires middleware, has query params for date range and/or song range
	e.DELETE("/tracks/history/:trackId", nil)          // Delete song from play history, Requires middleware
	e.GET("/tracks/artist/:artistId", nil)             // Get tracks by artist, has query params for limit and offset
	e.POST("/tracks/:trackId/like", nil)               // Like a track, Requires middleware, keep in mind this just adds to liked playlist
	e.DELETE("/tracks/:trackId/unlike", nil)           // Unlike a track, Requires middleware, keep in mind this just removes from liked playlist

	// -- Playlist Handlers
	e.GET("/playlists", nil)                                    // Get user's public and private playlists, Requires middleware
	e.GET("/playlists/user/:userId", nil)                       // Get public playlists created by user
	e.POST("/playlists", h.InsertPlaylistHandler)               // Create a new playlist, Requires middleware
	e.GET("/playlists/:playlistId", nil)                        // Get specific playlist by its id
	e.PUT("/playlists/:playlistId", h.UpdatePlaylistHandler)    // Update a playlist by its id, Requires middleware
	e.DELETE("/playlists/:playlistId", h.DeletePlaylistHandler) // Delete a playlist by its id, Requires middleware
	e.GET("/playlists/:playlistId/tracks", nil)                 // Get songs in playlist by playlist id
	e.POST("/playlists/:playlistId/tracks", nil)                // Add song to playlist, Requires middleware
	e.DELETE("/playlists/:playlistId/tracks", nil)              // Remove song from playlist, Requires middleware
	e.GET("/playlists/new-releases", nil)                       // Get new releases, has query params for limit, max return is 20

	// -- Following other users Handlers
	e.GET("/users/:userId/following", nil)   // Get users that user is following
	e.POST("/users/:userId/follow", nil)     // Follow a user, Requires middleware
	e.DELETE("/users/:userId/unfollow", nil) // Unfollow a user, Requires middleware

	// -- Following other artists Handlers
	e.GET("/artists/:artistId/followers", nil)   // Get artists that user is following
	e.POST("/artists/:artistId/follow", nil)     // Follow an artist, Requires middleware
	e.DELETE("/artists/:artistId/unfollow", nil) // Unfollow an artist, Requires middleware

	// -- Artist Handlers
	e.GET("/artists/:artistId", nil) // Get artist by id

	// -- Search Handlers
	e.GET("/search", nil) // Has query param ?q=<query>&type=<type>
}
