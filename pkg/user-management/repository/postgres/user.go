package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/shaream/pkg/models"

	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
	db sqlx.ExtContext
}

func NewUserRepository(db sqlx.ExtContext) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	q, args, err := sq.Insert("users").
		Columns(`
			id,
			email
		`).
		Values(
			user.ID,
			user.Email,
		).ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
