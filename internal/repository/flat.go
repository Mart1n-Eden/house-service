package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"house-service/internal/model"
)

func (r *repo) CreateFlat(ctx context.Context, houseId int, price int, rooms int) (flat *model.Flat, err error) {
	err = r.transact(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		flat, err = r.insertNewFlat(ctx, houseId, price, rooms)
		if err != nil {
			return err
		}

		err = r.updateLastDateHouse(ctx, houseId)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return flat, nil
}

func (r *repo) UpdateFlat(ctx context.Context, id int, status string) (*model.Flat, error) {
	query := `UPDATE flat SET status = $1 WHERE id = $2 RETURNING *`

	res := &model.Flat{}

	err := r.db.QueryRowxContext(ctx, query, status, id).
		Scan(&res.Id, &res.HouseId, &res.Price, &res.Rooms, &res.Status)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repo) updateLastDateHouse(ctx context.Context, id int) error {
	query := `UPDATE house SET update_at = now() WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) insertNewFlat(ctx context.Context, houseId int, price int, rooms int) (*model.Flat, error) {
	query := `INSERT INTO flat (house_id, price, rooms) VALUES ($1, $2, $3) RETURNING *`

	res := &model.Flat{}

	err := r.db.QueryRowxContext(ctx, query, houseId, price, rooms).
		Scan(&res.Id, &res.HouseId, &res.Price, &res.Rooms, &res.Status)
	if err != nil {
		return nil, err
	}

	return res, nil
}
