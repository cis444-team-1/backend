package repositories

import (
	"context"
	"database/sql"
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

func (r *TrackRepository) GetTracksByUserID(userID string) ([]*models.GetTrackResponse, error) {
	query := `
		SELECT track_id, title, image_src, audio_src, created_at, updated_at, uploaded_id, album_title, description, duration_seconds, artist_name
		FROM tracks WHERE uploaded_id = $1
	`

	ctx := context.Background()
	rows, err := r.db.QueryContext(ctx, query, userID)
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
			&track.UploadedId,
			&track.AlbumTitle,
			&track.Description,
			&track.DurationSeconds,
			&track.ArtistName,
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
	searchQuery := `
        select track_id, title, image_src, audio_src, created_at, updated_at, artist_name, album_title, description, duration_seconds, uploaded_id
        from tracks
        where search_vector @@ to_tsquery('english', $1)
        order by ts_rank(search_vector, to_tsquery('english', $1)) desc
        limit 50
	`

	rows, err := r.db.Query(searchQuery, query)
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
			&track.UploadedId,
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
		SELECT track_id, title, image_src, audio_src, created_at, updated_at, artist_name, album_title, description, duration_seconds, uploaded_id
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
		&track.UploadedId,
	)

	return track, err
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
			duration_seconds,
			uploaded_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning track_id;
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
		uploadedByID,
	)

	if err != nil {
		return "", err
	}

	return trackID.String(), nil
}

func (r *TrackRepository) AddTrackToPlayHistory(userID string, trackID string) error {
	query := `
		INSERT INTO play_history (
			play_id,
			user_id,
			track_id,
			played_at
		) VALUES (
			$1, $2, $3, $4
		) ON CONFLICT (user_id, track_id) DO UPDATE SET played_at = EXCLUDED.played_at;
	`

	_, err := r.db.Exec(query, uuid.New(), userID, trackID, time.Now())
	return err
}

func (r *TrackRepository) RemoveTrackFromPlayHistory(userID string, trackID string) error {
	query := `
		DELETE FROM play_history WHERE user_id = $1 AND track_id = $2;
	`
	_, err := r.db.Exec(query, userID, trackID)
	return err
}

func (r *TrackRepository) GetPlayHistory(userID string) ([]*models.GetTrackResponseWithPlayDate, error) {
	query := `
		SELECT
			tracks.track_id,
			tracks.title,
			tracks.image_src,
			tracks.audio_src,
			tracks.created_at,
			tracks.updated_at,
			tracks.artist_name,
			tracks.album_title,
			tracks.description,
			tracks.duration_seconds,
			tracks.uploaded_id,
			play_history.played_at
		FROM play_history
		JOIN tracks ON play_history.track_id = tracks.track_id
		WHERE play_history.user_id = $1
		ORDER BY play_history.played_at DESC;
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tracks []*models.GetTrackResponseWithPlayDate
	for rows.Next() {
		var track models.GetTrackResponseWithPlayDate
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
			&track.UploadedId,
			&track.PlayedAt,
		)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, &track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetTopCharts() ([]*models.GetTrackResponse, error) {
	query := `
		SELECT
			t.track_id,
			t.title,
			t.image_src,
			t.audio_src,
			t.created_at,
			t.updated_at,
			t.uploaded_id,
			t.album_title,
			t.description,
			t.duration_seconds
		FROM tracks t
		JOIN play_history ph ON t.track_id = ph.track_id
		GROUP BY t.track_id
		ORDER BY COUNT(ph.track_id) DESC
		LIMIT 25;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tracks []*models.GetTrackResponse

	for rows.Next() {
		var track models.GetTrackResponse
		if err := rows.Scan(
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
		); err != nil {
			return nil, err
		}
		tracks = append(tracks, &track)
	}

	return tracks, nil
}

func (r *TrackRepository) GetTrendingTracks() ([]*models.TrendingTrack, error) {
	query := `
		SELECT
			t.track_id,
			t.title,
			t.image_src,
			t.audio_src,
			t.created_at,
			t.updated_at,
			t.uploaded_id,
			t.album_title,
			t.description,
			t.duration_seconds,
			COUNT(ph.track_id) AS play_count
		FROM tracks t
		JOIN play_history ph ON t.track_id = ph.track_id
		WHERE ph.played_at >= NOW() - INTERVAL '7 days'
		GROUP BY t.track_id
		ORDER BY play_count DESC
		LIMIT 25;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trendingTracks []*models.TrendingTrack

	for rows.Next() {
		var track models.TrendingTrack
		if err := rows.Scan(
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
			&track.PlayCount,
		); err != nil {
			return nil, err
		}
		trendingTracks = append(trendingTracks, &track)
	}

	return trendingTracks, nil
}
