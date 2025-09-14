package service

import (
	"github.com/mykytaserdiuk/shaream/pkg/db"
	"github.com/mykytaserdiuk/shaream/pkg/db/postgres"
	"github.com/mykytaserdiuk/shaream/pkg/file-storage/repository"
)

func NewServices(db *postgres.DB, repoManager repository.RepoManager, minio db.S3Storage) *Services {
	return &Services{
		FileSvc: NewFileService(db, repoManager, minio),
	}
}
