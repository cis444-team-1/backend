package models

import (
	"time"

	"github.com/google/uuid"
)

type Track struct {
	TrackId         uuid.UUID `json:"track_id"`
	Title           string    `json:"title"`
	ImageSrc        string    `json:"image_src"`
	AudioSrc        string    `json:"audio_src"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	UploadedId      uuid.UUID `json:"uploaded_id"`
	ArtistId        uuid.UUID `json:"artist_id"`
	AlbumTitle      uuid.UUID `json:"album_title"`
	Description     string    `json:"description"`
	Lyrics          string    `json:"lyrics"`
	DurationSeconds int       `json:"duration_seconds"`
}
