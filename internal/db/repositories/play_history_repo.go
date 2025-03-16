package repositories

import "database/sql"

type PlayHistoryRepository struct {
	db *sql.DB
}

func NewPlayHistoryRepository(db *sql.DB) *PlayHistoryRepository {
	return &PlayHistoryRepository{
		db: db,
	}
}
