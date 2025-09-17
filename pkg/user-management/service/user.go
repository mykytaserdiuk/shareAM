package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"github.com/mykytaserdiuk/shaream/pkg/user-management/repository"

	"github.com/mykytaserdiuk/shaream/pkg/db/postgres"
	"github.com/mykytaserdiuk/shaream/pkg/models"
)

type UserService struct {
	db          *postgres.DB
	repoManager repository.RepoManager
}

func NewUserService(db *postgres.DB, repoManager repository.RepoManager) *UserService {
	return &UserService{db: db, repoManager: repoManager}
}

func (s *UserService) CreateUser(ctx context.Context, creds *models.Credentials) (*models.User, error) {
	tx, err := s.db.Begintx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	id := uuid.New().String()
	user := &models.User{
		ID:    id,
		Email: nil,
	}

	err = s.repoManager.NewUserRepo(tx).CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	creds.Name = id
	hash := sha256.New()
	hash.Write([]byte(creds.Password))

	creds.Hash = hex.EncodeToString(hash.Sum(nil))
	err = s.repoManager.NewCredentialsRepo(tx).InsertCredentials(ctx, creds)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return user, nil
}
