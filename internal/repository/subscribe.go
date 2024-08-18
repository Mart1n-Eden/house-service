package repository

import (
	"context"
	"fmt"
	"strings"

	"house-service/internal/domain"
	tools "house-service/pkg/utils/dbErrors"
)

func (r *Repo) NewSubscription(ctx context.Context, email string, houseId int) error {
	query := `INSERT INTO subscription (email, house_id) 
				VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, email, houseId)
	if err != nil {
		return tools.PrepareError(err)
	}

	return nil
}

func (r *Repo) GetMessagesForSubscription(ctx context.Context) ([]domain.Message, error) {
	list, err := r.getEmailsAndHouseIds(ctx)
	if err != nil {
		return nil, tools.PrepareError(err)
	}

	var messages []domain.Message

	for id, emails := range list {
		flats, err := r.getFlats(ctx, id)
		if err != nil {
			return nil, tools.PrepareError(err)
		}

		if len(flats) == 0 {
			continue
		}

		var msg strings.Builder
		msg.WriteString(fmt.Sprintf("New flats in house %d : ", id))

		for i := range flats {
			msg.WriteString(fmt.Sprintf("FlatId %d, Rooms %d, Price %d\n", flats[i].Id, flats[i].Rooms, flats[i].Price))
		}

		for i := range emails {
			messages = append(messages, domain.Message{
				Recipient: emails[i],
				Message:   msg.String(),
			})
		}

	}

	return messages, nil
}

func (r *Repo) getEmailsAndHouseIds(ctx context.Context) (map[int][]string, error) {
	query := `SELECT email, house_id FROM subscription ORDER BY house_id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, tools.PrepareError(err)
	}

	list := make(map[int][]string)

	for rows.Next() {
		var email string
		var houseId int
		err = rows.Scan(&email, &houseId)
		if err != nil {
			return nil, tools.PrepareError(err)
		}
		list[houseId] = append(list[houseId], email)
	}

	if err := rows.Err(); err != nil {
		return nil, tools.PrepareError(err)
	}

	return list, nil
}

func (r *Repo) getFlats(ctx context.Context, houseId int) ([]domain.Flat, error) {
	query := `SELECT id, house_id, price, rooms
				FROM flat 
				WHERE status = 'approved'
				AND house_id = $1
				AND updated_at > NOW() - INTERVAL '12 hours'`

	rows, err := r.db.QueryContext(ctx, query, houseId)
	if err != nil {
		return nil, tools.PrepareError(err)
	}

	var flats []domain.Flat
	for rows.Next() {
		var flat domain.Flat
		err = rows.Scan(&flat.Id, &flat.HouseId, &flat.Price, &flat.Rooms)
		if err != nil {
			return nil, tools.PrepareError(err)
		}
		flats = append(flats, flat)
	}

	if err := rows.Err(); err != nil {
		return nil, tools.PrepareError(err)
	}

	return flats, nil
}
