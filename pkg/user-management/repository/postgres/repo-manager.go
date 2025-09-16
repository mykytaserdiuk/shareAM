package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/shaream/pkg/user-management/repository"
)

type RepoManager struct{}

func NewRepoManager() *RepoManager {
	return &RepoManager{}
}

func (rm *RepoManager) NewUserRepo(db sqlx.ExtContext) repository.UserRepo {
	return NewUserRepository(db)
}
func (rm *RepoManager) NewCredentialsRepo(db sqlx.ExtContext) repository.CredentialsRepo {
	return NewCredentialsRepo(db)
}
