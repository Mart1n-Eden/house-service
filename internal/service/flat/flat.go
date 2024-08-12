package flat

import (
	"context"

	"github.com/jmoiron/sqlx"
	"house-service/internal/model"
)

type FlatRepo interface {
	CreateFlat(ctx context.Context, houseId int, price int, rooms int) (*model.Flat, error)
	UpdateFlat(ctx context.Context, id int, status string) (*model.Flat, error)
}

type Transactor interface {
	Transact(ctx context.Context, f func(ctx context.Context, tx *sqlx.Tx) error) error
}

type Service struct {
	repo       FlatRepo
	transactor Transactor
}

func New(repo FlatRepo) *Service {
	return &Service{
		repo: repo,
		//transactor: transactor,
	}
}

func (s *Service) CreateFlat(ctx context.Context, houseId int, price int, rooms int) (*model.Flat, error) {
	// TODO: add validation
	return s.repo.CreateFlat(ctx, houseId, price, rooms)
}

func (s *Service) UpdateFlat(ctx context.Context, id int, status string) (*model.Flat, error) {
	// TODO: add validation
	return s.repo.UpdateFlat(ctx, 0, "")
}
