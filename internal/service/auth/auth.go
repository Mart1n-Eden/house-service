package auth

import (
	"context"
)

type AuthRepo interface {
	CreateUser(ctx context.Context, email string, password string, userType string) (string, error)
	Login(ctx context.Context, id string, password string) (string, error)
}

type JWTToken interface {
	CreateToken(role string) (string, error)
	ParseToken(header string) (string, error)
}

type Service struct {
	repo      AuthRepo
	tokenizer JWTToken
}

func New(repo AuthRepo, token JWTToken) *Service {
	return &Service{
		repo:      repo,
		tokenizer: token,
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

	return s.tokenizer.CreateToken(userType)
}

// TODO: delete ctx (?)
func (s *Service) DummyLogin(ctx context.Context, userType string) (string, error) {
	return s.tokenizer.CreateToken(userType)
}

func (s *Service) ParseToken(header string) (string, error) {
	return s.tokenizer.ParseToken(header)
}
