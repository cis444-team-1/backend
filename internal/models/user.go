package models

import (
	"time"

	"github.com/google/uuid"
)

// User metadata object
type UserProfile struct {
	UserId    uuid.UUID `json:"user_id"`
	Bio       string    `json:"bio"`
	ImageSrc  string    `json:"image_src"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

// User Identity Object
// https://supabase.com/docs/guides/auth/identity
type UserIdentity struct {
	ProviderId   string `json:"provider_id"`
	UserId       string `json:"user_id"`
	IdentityData string `json:"identity_data"`
	Id           string `json:"id"`
	Provider     string `json:"provider"`
	Email        string `json:"email"`
	CreatedAt    string `json:"created_at"`
	LastSignInAt string `json:"last_sign_in_at"`
	UpdatedAt    string `json:"updated_at"`
}

// User model from supabase authentication
// https://supabase.com/docs/guides/auth/users
type User struct {
	Id               string       `json:"id"`
	Aud              string       `json:"aud"`
	Role             string       `json:"role"`
	Email            string       `json:"email"`
	EmailConfirmedAt string       `json:"email_confirmed_at"`
	Phone            string       `json:"phone"`
	PhoneConfirmedAt string       `json:"phone_confirmed_at"`
	ConfirmedAt      string       `json:"confirmed_at"`
	LastSignInAt     string       `json:"last_sign_in_at"`
	CreatedAt        string       `json:"created_at"`
	UpdatedAt        string       `json:"updated_at"`
	IsAnonymous      bool         `json:"is_anonymous"`
	UserIdentity     UserIdentity `json:"user_identity"`
}

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

type UserLikesTracks struct {
	UserId    uuid.UUID `json:"user_id"`
	TrackId   uuid.UUID `json:"track_id"`
	CreatedAt time.Time `json:"created_at"`
}
