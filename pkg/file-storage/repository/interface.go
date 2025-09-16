package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/shaream/pkg/models"
)

type RepoManager interface {
	NewFileRepo(db sqlx.ExtContext) FileRepo
}

type FileRepo interface {
	InsertFiles(ctx context.Context, files []*models.File) error
	GetFileByID(ctx context.Context, id string) (*models.File, error)
}
