package service

import (
	"context"
	"github.com/mykytaserdiuk/shaream/pkg/db"
	"github.com/mykytaserdiuk/shaream/pkg/models"
	"mime/multipart"
	"time"
)

type FileService struct {
	minio *db.MinioDB
}

func NewFileService(db *db.MinioDB) *FileService {
	return &FileService{minio: db}
}

func (s *FileService) Upload(ctx context.Context, token *string, file multipart.File, header *multipart.FileHeader) (string, error) {
	fileModel := &models.File{
		Name:      header.Filename + "_" + time.Now().Format("2006-01-02_15:04:05"),
		UserToken: token,
		Size:      header.Size,
	}

	err := s.minio.Upload(ctx, fileModel, file)
	if err != nil {
		return "", err
	}

	url := s.minio.GetLinkFromFile(ctx, fileModel)

	return url, nil
}
