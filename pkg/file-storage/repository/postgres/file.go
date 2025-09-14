package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/shaream/pkg/models"
)

type FileRepository struct {
	db sqlx.ExtContext
}

func NewFileRepository(db sqlx.ExtContext) *FileRepository {
	return &FileRepository{db: db}
}

func (FileRepository) InsertFiles(ctx context.Context, files []*models.File) error {
	//TODO implement me
	panic("implement me")
}
