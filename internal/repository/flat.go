package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"house-service/internal/domain"
	tools "house-service/pkg/utils/dbErrors"
)

func (r *Repo) CreateFlat(ctx context.Context, houseId int, price int, rooms int) (flat *domain.Flat, err error) {
	err = r.transact(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		flat, err = r.insertNewFlat(ctx, houseId, price, rooms)
		if err != nil {
			return tools.PrepareError(err)
		}

		err = r.updateLastDateHouse(ctx, houseId)
		if err != nil {
			return tools.PrepareError(err)
		}

		return nil
	})

	if err != nil {
		return nil, tools.PrepareError(err)
	}

	return flat, nil
}

func (r *Repo) UpdateFlat(ctx context.Context, id int, status string) (*domain.Flat, error) {
	query := `SELECT status, updated_by FROM flat WHERE id = $1`

	var currentStatus string
	var updatedBy *string
	if err := r.db.QueryRowxContext(ctx, query, id).Scan(&currentStatus, &updatedBy); err != nil {
		fmt.Println(err.Error())
		return nil, tools.PrepareError(err)
	}

	moderatorId, ok := ctx.Value("id").(string)
	if !ok {
		return nil, errors.New("moderator not found")
	}
	if updatedBy != nil {
		if currentStatus == "on_moderation" && *updatedBy != moderatorId {
			return nil, errors.New("flat already on moderation and not updated by this moderator")
		}
	}

	if (status == "approved" || status == "declined") && currentStatus != "on_moderation" {
		return nil, errors.New("flat not on moderation")
	}

	if status == "on_moderation" && currentStatus != "created" {
		return nil, errors.New("flat already passed moderation")
	}

	query = `UPDATE flat SET status = $1, updated_by = $2, updated_at = now() WHERE id = $3 RETURNING id, house_id, price, rooms, status`

	res := &domain.Flat{}

	if err := r.db.QueryRowxContext(ctx, query, status, moderatorId, id).
		Scan(&res.Id, &res.HouseId, &res.Price, &res.Rooms, &res.Status); err != nil {
		return nil, tools.PrepareError(err)
	}

	return res, nil
}
func (r *Repo) updateLastDateHouse(ctx context.Context, id int) error {
	query := `UPDATE house SET updated_at = now() WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return tools.PrepareError(err)
	}

	return nil
}

func (r *Repo) insertNewFlat(ctx context.Context, houseId int, price int, rooms int) (*domain.Flat, error) {
	query := `INSERT INTO flat (house_id, price, rooms) VALUES ($1, $2, $3) RETURNING id, house_id, price, rooms, status`

	res := &domain.Flat{}

	err := r.db.QueryRowxContext(ctx, query, houseId, price, rooms).
		Scan(&res.Id, &res.HouseId, &res.Price, &res.Rooms, &res.Status)
	if err != nil {
		return nil, tools.PrepareError(err)
	}

	return res, nil
}
