package flat

import (
	"context"
	"errors"
	"log/slog"
	"strconv"

	"github.com/jmoiron/sqlx"
	"house-service/internal/domain"
	"house-service/internal/logger"
	"house-service/pkg/utils/dbErrors"
)

type FlatRepo interface {
	CreateFlat(ctx context.Context, houseId int, price int, rooms int) (*domain.Flat, error)
	UpdateFlat(ctx context.Context, id int, status string) (*domain.Flat, error)
	GetHouse(ctx context.Context, id int) ([]domain.Flat, error)
}

type Cache interface {
	PutHouse(id string, house []domain.Flat) error
	GetHouse(id string) ([]domain.Flat, bool)
	Delete(id string)
}

type Transactor interface {
	Transact(ctx context.Context, f func(ctx context.Context, tx *sqlx.Tx) error) error
}

type Service struct {
	repo  FlatRepo
	cache Cache
}

func New(repo FlatRepo, cache Cache) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
	}
}

func (s *Service) CreateFlat(ctx context.Context, houseId int, price int, rooms int) (*domain.Flat, error) {
	return s.repo.CreateFlat(ctx, houseId, price, rooms)
}

func (s *Service) UpdateFlat(ctx context.Context, id int, status string) (*domain.Flat, error) {
	flat, err := s.repo.UpdateFlat(ctx, id, status)
	if err != nil {
		if err.Error() == dbErrors.ErrNotFound {
			return nil, errors.New("flat not found")
		}
		return nil, err
	}

	if flat.Status == "approved" {
		s.cache.Delete(strconv.Itoa(id))
	}

	return flat, nil
}

func (s *Service) GetHouse(ctx context.Context, id int) (house []domain.Flat, err error) {
	if ctx.Value("role") == "moderator" {
		return s.repo.GetHouse(ctx, id)
	}

	idStr := strconv.Itoa(id)

	house, ok := s.cache.GetHouse(idStr)
	if ok {
		return house, nil
	}

	house, err = s.repo.GetHouse(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(house) != 0 {
		err = s.cache.PutHouse(idStr, house)
		if err != nil {
			logger.Warn("cache error", slog.String("op", "flatService.GetHouse"), slog.String("error", err.Error()))
		}
	}

	return house, nil
}
