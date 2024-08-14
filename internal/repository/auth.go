package repository

import (
	"context"

	"house-service/pkg/utils/dbErrors"
)

func (r *repo) CreateUser(ctx context.Context, email string, password string, userType string) (userId string, err error) {
	query := `INSERT INTO users (email, password, user_type) VALUES ($1, $2, $3) RETURNING id`

	err = r.db.QueryRowxContext(ctx, query, email, password, userType).Scan(&userId)
	if err != nil {
		return "", dbErrors.PrepareError(err)
	}

	return userId, nil
}

func (r *repo) Login(ctx context.Context, id string, password string) (userType string, err error) {
	query := `SELECT user_type FROM users WHERE id = $1 AND password = $2`

	err = r.db.QueryRowxContext(ctx, query, id, password).Scan(&userType)
	if err != nil {
		return "", dbErrors.PrepareError(err)
	}

	return userType, nil
}

// func Transact(ctx context.Context, tx *sqlx.Tx, f func(tx *sqlx.Tx) error)
