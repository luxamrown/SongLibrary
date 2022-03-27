package usecase

import (
	"enigmacamp.com/song/model"
	"enigmacamp.com/song/repository"
)

type ShowSongListUseCase interface {
	ShowAll() []model.Song
}

type showSongListUseCase struct {
	repo repository.SongRepo
}

func (a *showSongListUseCase) ShowAll() []model.Song {
	return a.repo.GetAll()
}

func NewShowSongListUseCase(repo repository.SongRepo) ShowSongListUseCase {
	return &showSongListUseCase{
		repo: repo,
	}
}
