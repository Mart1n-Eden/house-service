package dbErrors

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

const (
	ErrNotFound         = "entity not found"
	ErrAlreadyExists    = "entity already exists"
	ErrFailedConnection = "failed connection"
)

func PrepareError(err error) error {
	pErr, ok := err.(*pq.Error)
	if !ok {
		switch err {
		case sql.ErrNoRows:
			return errors.New(ErrNotFound)
		default:
			return err
		}
	}

	switch pErr.Code {
	case "23503":
		return errors.New(ErrAlreadyExists)
	case "08006":
		return errors.New(ErrFailedConnection)
	}

	// TODO: add other dbErrors

	return err
}
