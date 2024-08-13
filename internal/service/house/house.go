package house

import (
	"context"
	"strconv"

	"house-service/internal/model"
)

type HouseRepo interface {
	CreateHouse(ctx context.Context, address string, year int, dev string) (*model.House, error)
	GetHouse(ctx context.Context, id int) ([]model.Flat, error)
}

type Cache interface {
	PutHouse(id string, house []model.Flat) error
	GetHouse(id string) ([]model.Flat, bool)
}

type Service struct {
	repo  HouseRepo
	cache Cache
}

func New(repo HouseRepo, cache Cache) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
	}
}

func (s *Service) CreateHouse(ctx context.Context, address string, year int, dev string) (*model.House, error) {
	// TODO: add validation
	return s.repo.CreateHouse(ctx, address, year, dev)
}

func (s *Service) GetHouse(ctx context.Context, id int) (house []model.Flat, err error) {
	// TODO: add validation
	idStr := strconv.Itoa(id)

	house, ok := s.cache.GetHouse(idStr)
	if ok {
		// TODO: update cache
		//s.cache.PutHouse(idStr, house)
		return house, nil
	}

	house, err = s.repo.GetHouse(ctx, id)
	if err != nil {
		// TODO: handling error
		return nil, err
	}

	// TODO: handling error
	s.cache.PutHouse(idStr, house)

	return house, nil
}
