package house

import (
	"context"

	"house-service/internal/domain"
)

type HouseRepo interface {
	CreateHouse(ctx context.Context, address string, year int, dev string) (*domain.House, error)
}

type Service struct {
	repo HouseRepo
}

func New(repo HouseRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateHouse(ctx context.Context, address string, year int, dev string) (*domain.House, error) {
	// TODO: add validation
	return s.repo.CreateHouse(ctx, address, year, dev)
}
