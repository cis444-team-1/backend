package models

import (
	"time"

	"github.com/google/uuid"
)

type PlayHistory struct {
	PlayId    uuid.UUID `json:"play_id"`
	UserId    uuid.UUID `json:"user_id"`
	TrackId   uuid.UUID `json:"track_id"`
	PlayedAt  time.Time `json:"played_at"`
	CreatedAt time.Time `json:"created_at"`
}

type PostPlayHistoryRequest struct {
	TrackId uuid.UUID `json:"track_id"`
}

type GetPlayHistoryResponse struct {
	PlayId    uuid.UUID `json:"play_id"`
	UserId    uuid.UUID `json:"user_id"`
	TrackId   uuid.UUID `json:"track_id"`
	PlayedAt  time.Time `json:"played_at"`
	CreatedAt time.Time `json:"created_at"`
}
