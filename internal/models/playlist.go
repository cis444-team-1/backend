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
	ImageSrc    string    `json:"image_src"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetPlaylistResponse struct {
	PlaylistId  uuid.UUID `json:"playlist_id"`
	UserId      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"is_public"`
	ImageSrc    string    `json:"image_src"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PostPlaylistRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`
	ImageSrc    string `json:"image_src"`
}

type PutPlaylistRequest struct {
	PlaylistId  uuid.UUID `json:"playlist_id"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	IsPublic    *bool     `json:"is_public"`
	ImageSrc    *string   `json:"image_src"`
}

type DeletePlaylistRequest struct {
	PlaylistId uuid.UUID `json:"playlist_id"`
}

type AddTrackToPlaylistRequest struct {
	TrackId  uuid.UUID `json:"track_id"`
	Position int       `json:"position"`
	AddedAt  time.Time `json:"added_at"`
}

type RemoveTrackFromPlaylistRequest struct {
	TrackId uuid.UUID `json:"track_id"`
}

type PlaylistUser struct {
	PlaylistId uuid.UUID `json:"playlist_id"`
	UserId     uuid.UUID `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
}
