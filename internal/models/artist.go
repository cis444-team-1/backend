package models

import (
	"time"

	"github.com/google/uuid"
)

type Artist struct {
	ArtistId    uuid.UUID `json:"artist_id"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
	ImageSrc    string    `json:"image_src"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TrackArtist struct {
	TrackId  uuid.UUID `json:"track_id"`
	ArtistId uuid.UUID `json:"artist_id"`
}
