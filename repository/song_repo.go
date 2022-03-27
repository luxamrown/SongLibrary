package repository

import "enigmacamp.com/song/model"

type SongRepo interface {
	Insert(newSong model.Song)
	GetSongByArtist(songArtist string) []model.Song
	GetAll() []model.Song
}
