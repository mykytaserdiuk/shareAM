package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/shaream/pkg/file-storage/repository"
)

type RepoManager struct{}

func NewRepoManager() repository.RepoManager {
	return &RepoManager{}
}

func (*RepoManager) NewFileRepo(db sqlx.ExtContext) repository.FileRepo {
	return NewFileRepository(db)
}
