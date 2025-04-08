package models

import (
	"time"

	"github.com/google/uuid"
)

// User model is already given by supabase library, please refer to supabase docs for user details
// https://supabase.com/docs/guides/auth/users

type UserFollowArtists struct {
	UserId    uuid.UUID `json:"user_id"`
	ArtistId  uuid.UUID `json:"artist_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserFollowsUsers struct {
	UserId         uuid.UUID `json:"user_id"`
	FollowedUserId uuid.UUID `json:"followed_user_id"`
	CreatedAt      time.Time `json:"created_at"`
}
