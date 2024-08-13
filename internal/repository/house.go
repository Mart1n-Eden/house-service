package repository

import (
	"context"

	"house-service/internal/model"
)

func (r *repo) CreateHouse(ctx context.Context, address string, year int, dev string) (*model.House, error) {
	query := `INSERT INTO house (address, year_built, developer) VALUES ($1, $2, $3) RETURNING *`

	res := &model.House{}

	err := r.db.QueryRowxContext(ctx, query, address, year, dev).
		Scan(&res.Id, &res.Address, &res.Year, &res.Developer, &res.CreatedAt, &res.UpdateAt)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repo) GetHouse(ctx context.Context, id int) ([]model.Flat, error) {
	query := `SELECT * FROM flat WHERE house_id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	//defer rows.Close() // TODO: error handling

	var flats []model.Flat

	for rows.Next() {
		var flat model.Flat
		err = rows.Scan(&flat.Id, &flat.HouseId, &flat.Price, &flat.Rooms, &flat.Status)
		if err != nil {
			return nil, err
		}
		flats = append(flats, flat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return flats, nil
}
