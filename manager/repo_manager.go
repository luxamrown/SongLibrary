package manager

import "enigmacamp.com/song/repository"

type RepoManager interface {
	SongRepo() repository.SongRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) SongRepo() repository.SongRepo {
	return repository.NewSongRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
