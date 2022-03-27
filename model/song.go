package model

type Song struct {
	SongArtist string `db:"song_artist"`
	SongAlbum  string `db:"song_album"`
	SongTitle  string `db:"song_title"`
}

func (s *Song) GetArtist() string {
	return s.SongArtist
}

func (s *Song) GetAlbum() string {
	return s.SongAlbum
}

func (s *Song) GetTitle() string {
	return s.SongTitle
}

func (s *Song) SetArtist(songArtist string) {
	s.SongArtist = songArtist
}

func (s *Song) SetAlbum(songAlbum string) {
	s.SongAlbum = songAlbum
}

func (s *Song) SetTitle(songTitle string) {
	s.SongTitle = songTitle
}

func NewSong(songArtist string, songAlbum string, songTitle string) Song {
	return Song{
		SongArtist: songArtist,
		SongAlbum:  songAlbum,
		SongTitle:  songTitle,
	}
}
