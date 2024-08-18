package handler

import (
	"context"
	"log/slog"

	"house-service/internal/domain"
)

type HouseService interface {
	CreateHouse(ctx context.Context, address string, year int, dev string) (*domain.House, error)
}

type FlatService interface {
	CreateFlat(ctx context.Context, houseId int, price int, rooms int) (*domain.Flat, error)
	UpdateFlat(ctx context.Context, id int, status string) (*domain.Flat, error)
	GetHouse(ctx context.Context, id int) ([]domain.Flat, error)
}

type AuthService interface {
	CreateUser(ctx context.Context, email string, password string, userType string) (string, error)
	Login(ctx context.Context, id string, password string) (string, error)
	DummyLogin(userType string) (string, error)
	ParseToken(header string) (string, string, error)
}

type Handler struct {
	log          *slog.Logger
	houseService HouseService
	flatService  FlatService
	authService  AuthService
}

func New(log *slog.Logger, house HouseService, flat FlatService, auth AuthService) *Handler {
	return &Handler{
		log:          log,
		houseService: house,
		flatService:  flat,
		authService:  auth,
	}
}
