package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cis444-team-1/backend/internal/models"
	"github.com/google/uuid"
)

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

func (r *TrackRepository) GetTracksByUserID(userID string) ([]*models.Track, error) {
	query := `
		SELECT * FROM tracks WHERE uploaded_id = $1
	`

	ctx := context.Background()
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tracks []*models.Track
	for rows.Next() {
		var track models.Track
		err := rows.Scan(
			&track.TrackId,
			&track.Title,
			&track.ImageSrc,
			&track.AudioSrc,
			&track.CreatedAt,
			&track.UpdatedAt,
			&track.UploadedId,
			&track.AlbumTitle,
			&track.Description,
			&track.DurationSeconds,
		)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, &track)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tracks, nil
}

func (r *TrackRepository) GetTracksBySearchQuery(query string) ([]*models.GetTrackResponse, error) {
	ctx := context.Background()
	// Optional: prefix match each term
	tsQuery := fmt.Sprintf("%s:*", query)

	rows, err := r.db.QueryContext(ctx, `
        select track_id, title, image_src, audio_src, created_at, updated_at, artist_name, album_title, description, duration_seconds
        from tracks
        where search_vector @@ to_tsquery('english', $1)
        order by ts_rank(search_vector, to_tsquery('english', $1)) desc
        limit 50
    `, tsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tracks []*models.GetTrackResponse
	for rows.Next() {
		var track models.GetTrackResponse
		err := rows.Scan(
			&track.TrackId,
			&track.Title,
			&track.ImageSrc,
			&track.AudioSrc,
			&track.CreatedAt,
			&track.UpdatedAt,
			&track.ArtistName,
			&track.AlbumTitle,
			&track.Description,
			&track.DurationSeconds,
		)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, &track)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tracks, nil
}

func (r *TrackRepository) GetTrackByID(trackID string) (*models.GetTrackResponse, error) {
	query := `
		SELECT track_id, title, image_src, audio_src, created_at, updated_at, artist_name, album_title, description, duration_seconds
		FROM tracks WHERE track_id = $1
	`

	track := &models.GetTrackResponse{}

	response := r.db.QueryRow(query, trackID)

	err := response.Scan(
		&track.TrackId,
		&track.Title,
		&track.ImageSrc,
		&track.AudioSrc,
		&track.CreatedAt,
		&track.UpdatedAt,
		&track.ArtistName,
		&track.AlbumTitle,
		&track.Description,
		&track.DurationSeconds,
	)

	return track, err
}

func (r *TrackRepository) GetTracksByGenre(genre string) ([]string, error) {
	// TODO: implement
	return nil, nil
}

func (r *TrackRepository) UpdateTrack(trackID string, data map[string]interface{}) error {
	// TODO: implement
	return nil
}

func (r *TrackRepository) DeleteTrack(trackID uuid.UUID) error {
	query := `
		DELETE FROM tracks WHERE track_id = $1
	`
	_, err := r.db.Exec(query, trackID)
	if err != nil {
		return err
	}
	return nil
}

func (r *TrackRepository) CreateTrack(uploadedByID string, track models.PostTrackRequest) (string, error) {
	query := `
		INSERT INTO tracks (
			track_id,
			title,
			image_src,
			audio_src,
			created_at,
			updated_at,
			artist_name,
			album_title,
			description,
			duration_seconds
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning track_id;
	`

	var trackID = uuid.New()

	_, err := r.db.Exec(query,
		trackID,
		track.Title,
		track.ImageSrc,
		track.AudioSrc,
		time.Now(),
		time.Now(),
		track.ArtistName,
		track.AlbumTitle,
		track.Description,
		track.DurationSeconds,
	)

	if err != nil {
		return "", err
	}

	return trackID.String(), nil
}
