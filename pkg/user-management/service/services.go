package service

import (
	"github.com/mykytaserdiuk/shaream/pkg/db/postgres"
	"github.com/mykytaserdiuk/shaream/pkg/user-management/repository"
)

func NewServices(db *postgres.DB, repoManager repository.RepoManager) *Services {
	return &Services{
		UserSvc: NewUserService(db, repoManager),
	}
}
