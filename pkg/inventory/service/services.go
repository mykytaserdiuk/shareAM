package service

import "github.com/mykytaserdiuk/shaream/pkg/db"

func NewServices(minio *db.MinioDB) *Services {
	return &Services{
		FileSvc: NewFileService(minio),
	}
}
