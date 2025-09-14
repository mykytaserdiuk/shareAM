package postgres

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func NewDB(connString string) (*DB, error) {
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) Begintx(ctx context.Context) (*sqlx.Tx, error) {
	tx, err := db.BeginTxx(ctx, &sql.TxOptions{
		//Isolation: db.isolationLevel,
		ReadOnly: false,
	})
	if err != nil {
		return nil, err
	}

	return tx, nil
}
