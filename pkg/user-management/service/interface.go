package service

import (
	"context"
	"github.com/mykytaserdiuk/shaream/pkg/models"
)

type Services struct {
	UserSvc UserSvc
}

type UserSvc interface {
	CreateUser(ctx context.Context, creds *models.Credentials) (*models.User, error)
}
