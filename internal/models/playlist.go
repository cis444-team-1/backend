package models

import (
	"time"

	"github.com/google/uuid"
)

type Playlist struct {
	PlaylistId  uuid.UUID `json:"playlist_id"`
	UserId      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"is_public"`
	CreatedAt   time.Time `json:"created_at"`
}

type PlaylistTrack struct {
	PlaylistId uuid.UUID `json:"playlist_id"`
	TrackId    uuid.UUID `json:"track_id"`
	Position   int       `json:"position"`
	AddedAt    time.Time `json:"added_at"`
}
