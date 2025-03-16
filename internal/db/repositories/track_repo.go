package repositories

import "database/sql"

type TrackRepository struct {
	db *sql.DB
}

func NewTrackRepository(db *sql.DB) *TrackRepository {
	return &TrackRepository{
		db: db,
	}
}

func (r *TrackRepository) GetTracksByPlaylistID(playlistID string) ([]string, error) {
	// TODO: implement
	return nil, nil
}

func (r *TrackRepository) GetTracksByAlbumID(albumID string) ([]string, error) {
	// TODO: implement
	return nil, nil
}

func (r *TrackRepository) GetTracksByArtistID(artistID string) ([]string, error) {
	// TODO: implement
	return nil, nil
}

func (r *TrackRepository) GetTracksByUserID(userID string) ([]string, error) {
	// TODO: implement
	return nil, nil
}

func (r *TrackRepository) GetTracksBySearchQuery(query string) ([]string, error) {
	// TODO: implement
	return nil, nil
}

func (r *TrackRepository) GetTrackByID(trackID string) (string, error) {
	// TODO: implement
	return "", nil
}

func (r *TrackRepository) GetTracksByGenre(genre string) ([]string, error) {
	// TODO: implement
	return nil, nil
}

func (r *TrackRepository) UpdateTrack(trackID string, data map[string]interface{}) error {
	// TODO: implement
	return nil
}

func (r *TrackRepository) DeleteTrack(trackID string) error {
	// TODO: implement
	return nil
}

func (r *TrackRepository) CreateTrack(data map[string]interface{}) (string, error) {
	// TODO: implement
	return "", nil
}
