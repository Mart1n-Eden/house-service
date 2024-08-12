package house

import (
	"context"

	"house-service/internal/model"
)

type HouseRepo interface {
	CreateHouse(ctx context.Context, address string, year int, dev string) (*model.House, error)
	GetHouse(ctx context.Context, id int) ([]model.Flat, error)
}

type Service struct {
	repo HouseRepo
}

func New(repo HouseRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateHouse(ctx context.Context, address string, year int, dev string) (*model.House, error) {
	// TODO: add validation
	return s.repo.CreateHouse(ctx, address, year, dev)
}

func (s *Service) GetHouse(ctx context.Context, id int) ([]model.Flat, error) {
	// TODO: add validation
	return s.repo.GetHouse(ctx, id)
}
