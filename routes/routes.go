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
	e.GET("/tracks/uploads", h.GetTracksByUserIDHandler, auth.AuthMiddleware)
	e.POST("/tracks", h.InsertTrackHandler, auth.AuthMiddleware)
	e.GET("/tracks/:trackId", h.GetTrackByTrackIDHandler)
	e.PUT("/tracks/:trackId", h.UpdateTrackHandler)         // Update a track by its id
	e.DELETE("/tracks/:trackId", h.DeleteTrackHandler)      // Delete a track by its id
	e.GET("/tracks/top-charts", h.GetTopChartsMusicHandler) // Get All-time popular songs
	e.GET("/tracks/trending", h.GetTopTrendingMusicHandler) // Get Recently popular songs
	//e.GET("/tracks/most-played", nil)                  // Get user's most played songs, Requires middleware
	e.POST("/tracks/history/:trackId", h.AddTrackToPlayHistoryHandler, auth.AuthMiddleware)        // Add song to play history, Requires middleware
	e.GET("/tracks/history", h.GetPlayHistoryHandler, auth.AuthMiddleware)                         // Get user's play history, Requires middleware, has query params for date range and/or song range
	e.DELETE("/tracks/history/:trackId", h.RemoveTrackFromPlayHistoryHandler, auth.AuthMiddleware) // Delete song from play history, Requires middleware
	//e.GET("/tracks/artist/:artistId", nil)             // Get tracks by artist, has query params for limit and offset
	//e.POST("/tracks/:trackId/like", nil)               // Like a track, Requires middleware, keep in mind this just adds to liked playlist
	//e.DELETE("/tracks/:trackId/unlike", nil)           // Unlike a track, Requires middleware, keep in mind this just removes from liked playlist

	// -- Playlist Handlers
	e.GET("/playlists", h.GetUsersPersonalPlaylistsHandler, auth.AuthMiddleware)
	e.GET("/playlists/user/:userId", nil) // Get public playlists created by user
	e.POST("/playlists", h.InsertPlaylistHandler, auth.AuthMiddleware)
	e.GET("/playlists/:playlistId", h.GetPlaylistHandler)
	e.PUT("/playlists/:playlistId", h.UpdatePlaylistHandler)    // Update a playlist by its id, Requires middleware
	e.DELETE("/playlists/:playlistId", h.DeletePlaylistHandler) // Delete a playlist by its id, Requires middleware
	e.GET("/playlists/:playlistId/tracks", h.GetTracksFromPlaylistHandler)
	e.POST("/playlists/:playlistId/tracks", h.AddTrackToPlaylistHandler)
	e.DELETE("/playlists/:playlistId/tracks", h.RemoveTrackFromPlaylistHandler) // Remove song from playlist, Requires middleware
	e.GET("/playlists/new-releases", h.GetNewReleasesHandler)

	// -- Search Handlers
	e.GET("/search", h.SearchHandler) // Has query param ?q=<query>&type=<type>

	e.GET("/users/:userId", h.GetUserHandler)         // DONE Get public user profile
	e.POST("/users/many", h.GetUsersByUserIDsHandler) // DONE
}
