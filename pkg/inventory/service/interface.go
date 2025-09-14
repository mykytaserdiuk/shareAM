package service

import (
	"context"
	"mime/multipart"
)

type Services struct {
	FileSvc FileSvc
}

type FileSvc interface {
	Upload(ctx context.Context, token *string, file multipart.File, header *multipart.FileHeader) (string, error)
}
