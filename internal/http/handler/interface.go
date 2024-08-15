package handler

import (
	"context"
	"log/slog"

	"house-service/internal/model"
)

type HouseService interface {
	CreateHouse(ctx context.Context, address string, year int, dev string) (*model.House, error)
	GetHouse(ctx context.Context, id int) ([]model.Flat, error)
}

type FlatService interface {
	CreateFlat(ctx context.Context, houseId int, price int, rooms int) (*model.Flat, error)
	UpdateFlat(ctx context.Context, id int, status string) (*model.Flat, error)
}

type AuthService interface {
	CreateUser(ctx context.Context, email string, password string, userType string) (string, error)
	Login(ctx context.Context, id string, password string) (string, error)
	DummyLogin(ctx context.Context, userType string) (string, error)
	ParseToken(header string) (string, error)
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
