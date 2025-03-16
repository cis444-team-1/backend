package models

import (
	"time"

	"github.com/google/uuid"
)

type Album struct {
	AlbumId     uuid.UUID `json:"album_id"`
	Title       string    `json:"title"`
	ArtistId    uuid.UUID `json:"artist_id"`
	ImageSrc    string    `json:"image_src"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
}
