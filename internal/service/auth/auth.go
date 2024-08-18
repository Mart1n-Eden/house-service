package auth

import (
	"context"

	"github.com/google/uuid"
	"house-service/internal/domain"
)

type AuthRepo interface {
	CreateUser(ctx context.Context, email string, password string, userType string) (string, error)
	Login(ctx context.Context, id string, password string) (domain.User, error)
}

type JWTToken interface {
	CreateToken(user domain.User) (string, error)
	ParseToken(header string) (string, string, error)
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
	user, err := s.repo.Login(ctx, id, password)
	if err != nil {
		return "", err
	}

	return s.tokenizer.CreateToken(user)
}

func (s *Service) DummyLogin(userType string) (string, error) {
	userID := uuid.New().String()

	user := domain.User{
		Id:       userID,
		UserType: userType,
	}

	return s.tokenizer.CreateToken(user)
}

func (s *Service) ParseToken(header string) (string, string, error) {
	return s.tokenizer.ParseToken(header)
}
