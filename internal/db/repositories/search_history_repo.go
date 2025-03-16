package repositories

import "database/sql"

type SearchHistoryRepository struct {
	db *sql.DB
}

func NewSearchHistoryRepository(db *sql.DB) *SearchHistoryRepository {
	return &SearchHistoryRepository{
		db: db,
	}
}
