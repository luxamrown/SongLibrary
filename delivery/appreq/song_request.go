package appreq

type SongRequest struct {
	SongArtist string `json:"song_artist"`
	SongAlbum  string `json:"song_album"`
	SongTitle  string `json:"song_title"`
}
