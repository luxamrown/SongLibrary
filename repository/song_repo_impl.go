package repository

import (
	"enigmacamp.com/song/model"
	"github.com/jmoiron/sqlx"
)

type SongRepoImpl struct {
	songDb *sqlx.DB
}

func (s *SongRepoImpl) Insert(newSong model.Song) (model.Song, error) {
	_, err := s.songDb.Exec("INSERT INTO song_list(song_artist, song_album, song_title) VALUES($1, $2, $3)", newSong.GetArtist(), newSong.GetAlbum(), newSong.GetTitle())
	if err != nil {
		return model.Song{}, err
	}
	return newSong, nil
}

func (s *SongRepoImpl) GetSongByArtist(songArtist string) ([]model.Song, error) {
	var selectedSong []model.Song
	err := s.songDb.Select(&selectedSong, "SELECT * FROM song_list WHERE song_artist = $1", songArtist)
	if err != nil {
		return []model.Song{}, err
	}
	return selectedSong, nil
}

func (s *SongRepoImpl) GetAll() ([]model.Song, error) {
	var songList []model.Song
	err := s.songDb.Select(&songList, "SELECT * FROM song_list")
	if err != nil {
		return []model.Song{}, err
	}
	// s.songDb.Get(&songList, "SELECT * FROM song_list")
	return songList, nil
}

func NewSongRepo(songDb *sqlx.DB) SongRepo {
	songRepo := SongRepoImpl{
		songDb: songDb,
	}
	return &songRepo
}
