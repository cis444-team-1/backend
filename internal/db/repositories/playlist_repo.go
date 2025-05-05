package repositories

import (
	"database/sql"
	"time"

	"github.com/cis444-team-1/backend/internal/models"
	"github.com/google/uuid"
)

type PlaylistRepository struct {
	db *sql.DB
}

func NewPlaylistRepository(db *sql.DB) *PlaylistRepository {
	return &PlaylistRepository{
		db: db,
	}
}

func (r *PlaylistRepository) GetPlaylistByID(playlistID string) (*models.GetPlaylistResponse, error) {
	query := `
		SELECT playlist_id, user_id, title, description, is_public, image_src, created_at, updated_at
		FROM playlists
		WHERE playlist_id = $1;
	`
	response := r.db.QueryRow(query, playlistID)

	playlist := &models.GetPlaylistResponse{}

	err := response.Scan(
		&playlist.PlaylistId,
		&playlist.UserId,
		&playlist.Title,
		&playlist.Description,
		&playlist.IsPublic,
		&playlist.ImageSrc,
		&playlist.CreatedAt,
		&playlist.UpdatedAt,
	)

	return playlist, err
}

func (r *PlaylistRepository) GetPlaylistsByUserID(userID string, isPublic bool) ([]string, error) {
	// TODO: implement
	return nil, nil
}

func (r *PlaylistRepository) UpdatePlaylist(playlistID string, data map[string]interface{}) error {
	// TODO: implement
	return nil
}

func (r *PlaylistRepository) DeletePlaylist(playlistID string) error {
	query := `
		DELETE FROM playlists WHERE playlist_id = $1
	`

	_, err := r.db.Exec(query, playlistID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PlaylistRepository) CreatePlaylist(uploadedByID string, playlist models.PostPlaylistRequest) (string, error) {
	query := `
		INSERT INTO playlists (
			playlist_id,
			title,
			description,
			is_public,
			image_src,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7 
		);`

	var playlistID = uuid.New()

	_, err := r.db.Exec(
		query,
		playlistID,
		playlist.Title,
		playlist.Description,
		playlist.IsPublic,
		playlist.ImageSrc,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return "", err
	}

	return playlistID.String(), nil
}

func (r *PlaylistRepository) AddTrackToPlaylist(playlistID string, trackID string, position int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	updatePositionQuery := `
		UPDATE playlist_tracks
		SET position = position + 1
		WHERE playlist_id = $1 AND position >= $2
	`

	query := `
		INSERT INTO playlist_tracks (
			playlist_id,
			track_id,
			position,
			added_at
		) VALUES (
			$1, $2, $3, $4
		);
	`

	_, err = tx.Exec(updatePositionQuery, playlistID, position)
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, playlistID, trackID, position, time.Now())
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PlaylistRepository) RemoveTrackFromPlaylist(playlistID string, trackID string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	getPositionQuery := `
		SELECT position
		FROM playlist_tracks
		WHERE playlist_id = $1 AND track_id = $2;
	`

	removeTrackQuery := `
		DELETE FROM playlist_tracks
		WHERE playlist_id = $1 AND track_id = $2;
	`

	updatePositionQuery := `
		UPDATE playlist_tracks
		SET position = position - 1
		WHERE playlist_id = $1 AND position > $2;
	`

	var pos int
	err = tx.QueryRow(getPositionQuery, playlistID, trackID).Scan(&pos)

	if err != nil {
		return err
	}

	_, err = tx.Exec(removeTrackQuery, playlistID, trackID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(updatePositionQuery, playlistID, pos)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PlaylistRepository) GetTracksByPlaylistID(playlistID string) ([]models.GetTrackResponse, error) {
	query := `
		SELECT playlist_tracks.track_id, tracks.title, tracks.artist_name, tracks.audio_src, tracks.album_title, tracks.duration_seconds, tracks.image_src, tracks.created_at,
			   tracks.updated_at, tracks.description
		FROM playlist_tracks
		JOIN tracks ON playlist_tracks.track_id = tracks.track_id
		WHERE playlist_tracks.playlist_id = $1
		ORDER BY playlist_tracks.position;
	`

	rows, err := r.db.Query(query, playlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tracks []models.GetTrackResponse
	for rows.Next() {
		var track models.GetTrackResponse

		err := rows.Scan(
			&track.TrackId,
			&track.Title,
			&track.ArtistName,
			&track.AudioSrc,
			&track.AlbumTitle,
			&track.DurationSeconds,
			&track.ImageSrc,
			&track.CreatedAt,
			&track.UpdatedAt,
			&track.Description,
		)

		if err != nil {
			return nil, err
		}

		tracks = append(tracks, track)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tracks, nil
}

func (r *PlaylistRepository) GetPlaylistsBySearchQuery(query string) ([]*models.GetPlaylistResponse, error) {
	searchQuery := `
        select playlist_id, user_id, title, description, is_public, image_src, created_at, updated_at
        from playlists
        where search_vector @@ to_tsquery('english', $1)
        order by ts_rank(search_vector, to_tsquery('english', $1)) desc
        limit 50
	`

	rows, err := r.db.Query(searchQuery, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var playlists []*models.GetPlaylistResponse
	for rows.Next() {
		var playlist models.GetPlaylistResponse
		err := rows.Scan(
			&playlist.PlaylistId,
			&playlist.UserId,
			&playlist.Title,
			&playlist.Description,
			&playlist.IsPublic,
			&playlist.ImageSrc,
			&playlist.CreatedAt,
			&playlist.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, &playlist)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return playlists, nil
}
