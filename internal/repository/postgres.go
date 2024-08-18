package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"house-service/internal/config"
)

type Repo struct {
	db *sqlx.DB
}

func NewConnection(ctx context.Context, cfg config.DB) (*sqlx.DB, error) {
	// TODO: add sslMode
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Name, cfg.Password)

	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func New(conn *sqlx.DB) *Repo {
	return &Repo{db: conn}
}
