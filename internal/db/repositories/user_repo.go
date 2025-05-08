package repositories

import (
	"database/sql"
	"encoding/json"

	"github.com/lib/pq"
	"github.com/nedpals/supabase-go"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByUserID(userID string) (*supabase.User, error) {

	query := `
		SELECT id, raw_user_meta_data FROM auth.users WHERE id = $1;
	`

	user := &supabase.User{}

	var metadata []uint8

	err := r.db.QueryRow(query, userID).Scan(
		&user.ID,
		&metadata,
	)

	if err != nil {
		return nil, err
	}

	// Convert metadata to json
	err = json.Unmarshal(metadata, &user.UserMetadata)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUsersByUserIDs(userIDs []string) ([]supabase.User, error) {
	query := `
		SELECT id, raw_user_meta_data FROM auth.users WHERE id = ANY($1);
	`

	var users []supabase.User

	rows, err := r.db.Query(query, pq.Array(userIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user supabase.User
		var metadata []uint8
		err := rows.Scan(
			&user.ID,
			&metadata,
		)
		if err != nil {
			return nil, err
		}
		// Convert metadata to json
		err = json.Unmarshal(metadata, &user.UserMetadata)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
