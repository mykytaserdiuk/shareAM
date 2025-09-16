package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/shaream/pkg/models"

	sq "github.com/Masterminds/squirrel"
)

type CredentialsRepo struct {
	db sqlx.ExtContext
}

func NewCredentialsRepo(db sqlx.ExtContext) *CredentialsRepo {
	return &CredentialsRepo{db: db}
}

func (r *CredentialsRepo) InsertCredentials(ctx context.Context, creds *models.Credentials) error {
	q, args, err := sq.Insert("credentials").
		Columns(`
			name, 
			hash		
		`).
		Values(
			creds.Name,
			creds.Hash,
		).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
