package house

import (
	"context"
	"errors"

	"house-service/internal/domain"
	"house-service/pkg/utils/dbErrors"
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
	house, err := s.repo.CreateHouse(ctx, address, year, dev)
	if err != nil {
		if err.Error() != dbErrors.ErrFailedConnection {
			return nil, errors.New("house already exists")
		}
		return nil, err
	}

	return house, nil
}
