-- Custom user metadata connected to auth table
CREATE TABLE user_profiles (
  user_id UUID PRIMARY KEY REFERENCES auth.users(id) ON DELETE CASCADE,
  bio TEXT,
  image_src TEXT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Playlists table
CREATE TABLE playlists (
  playlist_id UUID PRIMARY KEY,
  user_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
  title VARCHAR(100) NOT NULL,
  description TEXT,
  is_public BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Artists table (separate from just a string in tracks)
CREATE TABLE artists (
  artist_id UUID PRIMARY KEY,
  display_name VARCHAR(100) NOT NULL,
  bio TEXT,
  image_src TEXT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Albums table
CREATE TABLE albums (
  album_id UUID PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  artist_id UUID REFERENCES artists(artist_id) ON DELETE CASCADE,
  image_src TEXT,
  release_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tracks table
CREATE TABLE tracks (
  track_id UUID PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  image_src TEXT,
  audio_src TEXT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  uploaded_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
  artist_id UUID REFERENCES artists(artist_id) ON DELETE CASCADE,
  album_id UUID REFERENCES albums(album_id) ON DELETE CASCADE,
  description TEXT,
  lyrics TEXT,
  duration INTEGER NOT NULL -- duration in seconds
);

-- Playlist tracks junction table
CREATE TABLE playlist_tracks (
  playlist_id UUID REFERENCES playlists(playlist_id) ON DELETE CASCADE,
  track_id UUID REFERENCES tracks(track_id) ON DELETE CASCADE,
  position INTEGER NOT NULL,
  added_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (playlist_id, track_id)
);

-- Play history table
CREATE TABLE play_history (
  play_id UUID PRIMARY KEY,
  user_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
  track_id UUID REFERENCES tracks(track_id) ON DELETE CASCADE,
  played_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMP WITH TIME ZONE DEFAULT (CURRENT_TIMESTAMP + INTERVAL '2 years')
);

-- Genre table
CREATE TABLE genres (
  genre_id UUID PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE
);

-- Track-artist junction (for multiple artists per track)
CREATE TABLE track_artists (
  track_id UUID REFERENCES tracks(track_id) ON DELETE CASCADE,
  artist_id UUID REFERENCES artists(artist_id) ON DELETE CASCADE,
  PRIMARY KEY (track_id, artist_id)
);

-- Track-genre junction
CREATE TABLE track_genres (
  track_id UUID REFERENCES tracks(track_id) ON DELETE CASCADE,
  genre_id UUID REFERENCES genres(genre_id) ON DELETE CASCADE,
  PRIMARY KEY (track_id, genre_id)
);

-- User follows other users
CREATE TABLE user_follows_users (
  user_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
  followed_user_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, followed_user_id)
);

-- User follows artists
CREATE TABLE user_follows_artists (
  user_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
  artist_id UUID REFERENCES artists(artist_id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, artist_id)
);

-- User likes a song
CREATE TABLE user_likes (
  user_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
  track_id UUID REFERENCES tracks(track_id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, track_id)
);

-- User recommendations table
CREATE TABLE recommendations (
  recommendation_id UUID PRIMARY KEY,
  user_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
  track_id UUID REFERENCES tracks(track_id) ON DELETE CASCADE,
  reason TEXT,
  source TEXT, -- 'ai', 'similar_users', 'playlist_based', etc.
  score FLOAT NOT NULL, -- 0-1 for recommendation score
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMP WITH TIME ZONE DEFAULT (CURRENT_TIMESTAMP + INTERVAL '30 days')
);

-- Search history table
CREATE TABLE search_history (
  search_id UUID PRIMARY KEY,
  user_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
  query TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMP WITH TIME ZONE DEFAULT (CURRENT_TIMESTAMP + INTERVAL '30 days')
);

-- Indexes for faster queries
CREATE INDEX idx_tracks_artist_id ON tracks(artist_id);
CREATE INDEX idx_tracks_album_id ON tracks(album_id);
CREATE INDEX idx_playlists_user_id ON playlists(user_id);
CREATE INDEX idx_play_history_user_id ON play_history(user_id);
CREATE INDEX idx_track_title ON tracks(title);
CREATE INDEX idx_artist_display_name ON artists(display_name);
CREATE INDEX idx_album_title ON albums(title);
CREATE INDEX idx_genre_name ON genres(name);
CREATE INDEX idx_user_follows_users ON user_follows_users(user_id, followed_user_id);
CREATE INDEX idx_user_likes ON user_likes(user_id, track_id);