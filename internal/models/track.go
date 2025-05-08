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
	AlbumTitle      string    `json:"album_title"`
	Description     string    `json:"description"`
	DurationSeconds int       `json:"duration_seconds"`
}

type GetTrackResponse struct {
	TrackId         uuid.UUID `json:"track_id"`
	Title           string    `json:"title"`
	ImageSrc        string    `json:"image_src"`
	AudioSrc        string    `json:"audio_src"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ArtistName      string    `json:"artist_name"`
	AlbumTitle      string    `json:"album_title"`
	Description     string    `json:"description"`
	DurationSeconds int       `json:"duration_seconds"`
	UploadedId      uuid.UUID `json:"uploaded_id"`
}

type GetTrackResponseWithPlayDate struct {
	TrackId         uuid.UUID `json:"track_id"`
	Title           string    `json:"title"`
	ImageSrc        string    `json:"image_src"`
	AudioSrc        string    `json:"audio_src"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ArtistName      string    `json:"artist_name"`
	AlbumTitle      string    `json:"album_title"`
	Description     string    `json:"description"`
	DurationSeconds int       `json:"duration_seconds"`
	UploadedId      uuid.UUID `json:"uploaded_id"`
	PlayedAt        time.Time `json:"played_at"`
}

type PostTrackRequest struct {
	Title           string `json:"title"`
	ImageSrc        string `json:"image_src"`
	AudioSrc        string `json:"audio_src"`
	AlbumTitle      string `json:"album_title"`
	Description     string `json:"description"`
	DurationSeconds int    `json:"duration_seconds"`
	ArtistName      string `json:"artist_name"`
}

type PutTrackRequest struct {
	TrackId         uuid.UUID `json:"track_id"`
	Title           *string   `json:"title"`
	ImageSrc        *string   `json:"image_src"`
	AudioSrc        *string   `json:"audio_src"`
	AlbumTitle      *string   `json:"album_title"`
	Description     *string   `json:"description"`
	DurationSeconds *int      `json:"duration_seconds"`
	ArtistName      *string   `json:"artist_name"`
}

type TrendingTrack struct {
	Track
	PlayCount int `json:"play_count"`
}
