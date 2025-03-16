package models

import "github.com/google/uuid"

type TrackGenre struct {
	TrackId uuid.UUID `json:"track_id"`
	GenreId uuid.UUID `json:"genre_id"`
}

type Genre struct {
	GenreId uuid.UUID `json:"genre_id"`
	Name    string    `json:"name"`
}
