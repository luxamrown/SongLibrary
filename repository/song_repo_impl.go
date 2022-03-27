package repository

import (
	"fmt"

	"enigmacamp.com/song/model"
	"github.com/jmoiron/sqlx"
)

type SongRepoImpl struct {
	songDb *sqlx.DB
}

func (s *SongRepoImpl) Insert(newSong model.Song) {
	_, err := s.songDb.Exec("INSERT INTO song_list(song_artist, song_album, song_title) VALUES($1, $2, $3)", newSong.GetArtist(), newSong.GetAlbum(), newSong.GetTitle())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("New Song Has Been Saved to Database")
}

func (s *SongRepoImpl) GetSongByArtist(songArtist string) []model.Song {
	var selectedSong []model.Song
	s.songDb.Select(&selectedSong, "SELECT * FROM song_list WHERE song_artist = $1", songArtist)
	return selectedSong
}

func (s *SongRepoImpl) GetAll() []model.Song {
	var songList []model.Song
	sql := `SELECT * FROM song_list`
	s.songDb.Select(&songList, sql)
	// s.songDb.Get(&songList, "SELECT * FROM song_list")
	return songList
}

func NewSongRepo(songDb *sqlx.DB) SongRepo {
	songRepo := SongRepoImpl{
		songDb: songDb,
	}
	return &songRepo
}
