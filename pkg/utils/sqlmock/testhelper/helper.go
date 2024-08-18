package testhelper

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zhashkevych/go-sqlxmock"
)

type SetupDBMockOption func(testing.TB, sqlmock.Sqlmock)

func SetupDBMock(tb testing.TB, opts ...SetupDBMockOption) *sqlx.DB {
	tb.Helper()

	db, dbMock, err := sqlmock.Newx(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(tb, err)

	tb.Cleanup(func() {
		assert.NoError(tb, dbMock.ExpectationsWereMet())
	})

	for _, o := range opts {
		o(tb, dbMock)
	}

	return db
}
