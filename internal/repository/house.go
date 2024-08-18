package repository

import (
	"context"

	"house-service/internal/domain"
)

func (r *Repo) CreateHouse(ctx context.Context, address string, year int, dev string) (*domain.House, error) {
	query := `INSERT INTO house (address, year_built, developer) VALUES ($1, $2, $3) RETURNING *`

	res := &domain.House{}

	err := r.db.QueryRowxContext(ctx, query, address, year, dev).
		Scan(&res.Id, &res.Address, &res.Year, &res.Developer, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Repo) GetHouse(ctx context.Context, id int) ([]domain.Flat, error) {
	role := ctx.Value("role").(string)

	var query string

	switch role {
	case "moderator":
		query = `SELECT * FROM flat WHERE house_id = $1`
	case "client":
		query = `SELECT * FROM flat WHERE house_id = $1 AND status = 'approved'`
	}

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	var flats []domain.Flat

	for rows.Next() {
		var flat domain.Flat
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
