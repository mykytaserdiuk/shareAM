package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/shaream/pkg/models"

	sq "github.com/Masterminds/squirrel"
)

type FileRepository struct {
	db sqlx.ExtContext
}

func NewFileRepository(db sqlx.ExtContext) *FileRepository {
	return &FileRepository{db: db}
}

func (r *FileRepository) InsertFiles(ctx context.Context, files []*models.File) error {
	builder := sq.
		Insert("files").
		Columns(`
			id, 
			name, 
			bucket,
			size, 
			user_token,
			url
		`)
	for _, f := range files {
		builder = builder.Values(
			f.ID,
			f.Name,
			f.Bucket,
			f.Size,
			f.UserToken,
			f.URL,
		)
	}

	q, args, err := builder.
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, q, args...)

	return err
}

func (r *FileRepository) GetFileByID(ctx context.Context, id string) (*models.File, error) {
	q, args, err := sq.Select(`
			id, 
			name, 
			bucket,
			size, 
			user_token,
			url
		`).
		From("files").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var file models.File
	err = sqlx.SelectContext(ctx, r.db, &file, q, args...)
	if err != nil {
		return nil, err
	}

	return &file, nil
}
