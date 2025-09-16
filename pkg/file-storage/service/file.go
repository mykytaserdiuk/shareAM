package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/mykytaserdiuk/shaream/pkg/db"
	"github.com/mykytaserdiuk/shaream/pkg/db/postgres"
	"github.com/mykytaserdiuk/shaream/pkg/file-storage/repository"
	"github.com/mykytaserdiuk/shaream/pkg/models"
	"mime/multipart"
	"time"
)

type FileService struct {
	minio       db.S3Storage
	db          *postgres.DB
	repoManager repository.RepoManager
}

func NewFileService(db *postgres.DB, repoManager repository.RepoManager, minio db.S3Storage) *FileService {
	return &FileService{
		minio:       minio,
		db:          db,
		repoManager: repoManager,
	}
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

	url := s.minio.GetFileLink(ctx, fileModel)
	fileModel.URL = url

	newUUID := uuid.New()
	fileModel.ID = newUUID.String()

	err = s.repoManager.NewFileRepo(s.db).InsertFiles(ctx, []*models.File{fileModel})
	if err != nil {
		return "", err
	}

	return fileModel.URL, nil
}
