package usecase

import (
	"enigmacamp.com/song/model"
	"enigmacamp.com/song/repository"
)

type FindSongByArtist interface {
	FindSongByArtist(songArtist string) []model.Song
}

type findSongByArtist struct {
	repo repository.SongRepo
}

func (a *findSongByArtist) FindSongByArtist(songArtist string) []model.Song {
	if len(songArtist) == 0 {
		return a.repo.GetAll()
	}
	return a.repo.GetSongByArtist(songArtist)

}

func NewFindSongByArtist(repo repository.SongRepo) FindSongByArtist {
	return &findSongByArtist{
		repo: repo,
	}
}
