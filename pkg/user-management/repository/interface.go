package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/shaream/pkg/models"
)

type RepoManager interface {
	NewUserRepo(db sqlx.ExtContext) UserRepo
	NewCredentialsRepo(db sqlx.ExtContext) CredentialsRepo
}
type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
}

type CredentialsRepo interface {
	InsertCredentials(ctx context.Context, creds *models.Credentials) error
}
