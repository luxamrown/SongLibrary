package repository

import "enigmacamp.com/song/model"

type SongRepo interface {
	Insert(newSong model.Song) (model.Song, error)
	GetSongByArtist(songArtist string) ([]model.Song, error)
	GetAll() ([]model.Song, error)
}
