package manager

import "enigmacamp.com/song/usecase"

type UseCaseManager interface {
	AddSongUseCase() usecase.AddSongUseCase
	ShowSongListUseCase() usecase.ShowSongListUseCase
	FindSongByArtist() usecase.FindSongByArtist
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) AddSongUseCase() usecase.AddSongUseCase {
	return usecase.NewAddSongUseCase(u.repo.SongRepo())
}

func (u *useCaseManager) ShowSongListUseCase() usecase.ShowSongListUseCase {
	return usecase.NewShowSongListUseCase(u.repo.SongRepo())
}

func (u *useCaseManager) FindSongByArtist() usecase.FindSongByArtist {
	return usecase.NewFindSongByArtist(u.repo.SongRepo())
}

func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
