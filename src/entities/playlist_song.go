package entities

type PlaylistSong struct {
	ID         string `json:"id" validate:"required"`
	PlaylistID string `json:"playlist_id"`
	SongID     string `json:"song_id"`
}
