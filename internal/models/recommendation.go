package models

import (
	"time"

	"github.com/google/uuid"
)

type Recommendation struct {
	RecommendationId uuid.UUID `json:"recommendation_id"`
	UserId           uuid.UUID `json:"user_id"`
	TrackId          uuid.UUID `json:"track_id"`
	Reason           string    `json:"reason"`
	Source           string    `json:"source"`
	Score            float64   `json:"score"`
	CreatedAt        time.Time `json:"created_at"`
	ExpiresAt        time.Time `json:"expires_at"`
}
