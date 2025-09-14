package db

import (
	"context"
	"github.com/mykytaserdiuk/shaream/pkg/models"
	"io"
)

type S3Storage interface {
	Upload(ctx context.Context, file *models.File, reader io.Reader) error

	GetFileLink(ctx context.Context, file *models.File) string
	SetupBucket(ctx context.Context, bucket string) error
	BucketExists(ctx context.Context, bucket string) (bool, error)
	MakeBucket(ctx context.Context, bucket string) error
}
