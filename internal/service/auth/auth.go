package auth

import (
	"context"

	"house-service/internal/token"
)

type AuthRepo interface {
	CreateUser(ctx context.Context, email string, password string, userType string) (string, error)
	Login(ctx context.Context, id string, password string) (string, error)
}

type Service struct {
	repo   AuthRepo
	secret string
}

func New(repo AuthRepo, secret string) *Service {
	return &Service{
		repo:   repo,
		secret: secret,
	}
}

func (s *Service) CreateUser(ctx context.Context, email string, password string, userType string) (string, error) {
	return s.repo.CreateUser(ctx, email, password, userType)
}

func (s *Service) Login(ctx context.Context, id string, password string) (string, error) {
	userType, err := s.repo.Login(ctx, id, password)
	if err != nil {
		return "", err
	}

	return token.CreateJWTToken(userType, s.secret)
}

// TODO: delete ctx (?)
func (s *Service) DummyLogin(ctx context.Context, userType string) (string, error) {
	return token.CreateJWTToken(userType, s.secret)
}
