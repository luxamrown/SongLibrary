package usecase

import (
	"enigmacamp.com/song/model"
	"enigmacamp.com/song/repository"
)

type AddSongUseCase interface {
	Register(songArtist string, songAlbum string, songTitle string) (model.Song, error)
}

type addSongUseCase struct {
	repo repository.SongRepo
}

func (a *addSongUseCase) Register(songArtist string, songAlbum string, songTitle string) (model.Song, error) {
	newSong := model.NewSong(songArtist, songAlbum, songTitle)
	return a.repo.Insert(newSong)
}

func NewAddSongUseCase(repo repository.SongRepo) AddSongUseCase {
	return &addSongUseCase{
		repo: repo,
	}
}
