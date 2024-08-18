package repository

import (
	"context"

	"house-service/internal/domain"
	"house-service/pkg/utils/dbErrors"
)

func (r *Repo) CreateUser(ctx context.Context, email string, password string, userType string) (userId string, err error) {
	query := `INSERT INTO users (email, password, user_type) VALUES ($1, $2, $3) RETURNING id`

	err = r.db.QueryRowxContext(ctx, query, email, password, userType).Scan(&userId)
	if err != nil {
		return "", dbErrors.PrepareError(err)
	}

	return userId, nil
}

func (r *Repo) Login(ctx context.Context, id string, password string) (domain.User, error) {
	query := `SELECT id, user_type FROM users WHERE id = $1 AND password = $2`

	var userType domain.User
	err := r.db.QueryRowxContext(ctx, query, id, password).Scan(&userType.Id, &userType.UserType)
	if err != nil {
		return domain.User{}, dbErrors.PrepareError(err)
	}

	return userType, nil
}
