package repository

import (
	"context"

	"house-service/internal/model"
)

func (r *repo) CreateFlat(ctx context.Context, houseId int, price int, rooms int) (*model.Flat, error) {
	query := `INSERT INTO flat (house_id, price, rooms) VALUES ($1, $2, $3) RETURNING *`

	res := &model.Flat{}

	// TODO: transaction
	err := r.db.QueryRowxContext(ctx, query, houseId, price, rooms).
		Scan(&res.Id, &res.HouseId, &res.Price, &res.Rooms, &res.Status)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repo) UpdateFlat(ctx context.Context, id int, status string) (*model.Flat, error) {
	query := `UPDATE flat SET status = $1 WHERE id = $2 RETURNING *`

	res := &model.Flat{}

	// TODO: transaction
	err := r.db.QueryRowxContext(ctx, query, status, id).
		Scan(&res.Id, &res.HouseId, &res.Price, &res.Rooms, &res.Status)
	if err != nil {
		return nil, err
	}

	return res, nil
}
