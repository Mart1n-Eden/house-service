package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zhashkevych/go-sqlxmock"
	"house-service/internal/sender"
	"house-service/internal/service/subscribe"

	"house-service/internal/cache"
	"house-service/internal/domain"
	"house-service/internal/http/handler"
	"house-service/internal/http/handler/model/response"
	"house-service/internal/logger"
	"house-service/internal/repository"
	"house-service/internal/service/auth"
	"house-service/internal/service/flat"
	"house-service/internal/service/house"
	"house-service/internal/token"
	"house-service/pkg/utils/sqlmock/testhelper"
)

func TestGetHouse(t *testing.T) {
	logger.MustInit("local")

	origin := []domain.Flat{
		{
			Id:      1,
			HouseId: 1,
			Price:   1000,
			Rooms:   3,
			Status:  "created",
		},
		{
			Id:      2,
			HouseId: 1,
			Price:   2000,
			Rooms:   4,
			Status:  "on_moderation",
		},
		{
			Id:      3,
			HouseId: 1,
			Price:   3000,
			Rooms:   5,
			Status:  "approved",
		},
		{
			Id:      4,
			HouseId: 1,
			Price:   4000,
			Rooms:   6,
			Status:  "declined",
		},
	}

	t.Run("GetHouseModerator", func(t *testing.T) {
		pg := testhelper.SetupDBMock(t, setupGetHouseDBMock(t, 1, "moderator", origin))

		repo := repository.New(pg)
		c := cache.New()

		tok := token.New("")
		send := sender.New()

		houseService := house.New(repo)
		flatService := flat.New(repo, c)
		authService := auth.New(repo, tok)
		subService := subscribe.New(repo, send)
		hnd := handler.New(houseService, flatService, authService, subService)

		req, err := http.NewRequest("GET", "/house/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.WithValue(req.Context(), "role", "moderator")
		req = req.WithContext(ctx)

		w := httptest.NewRecorder()

		hnd.GetHouse(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d, got %d", http.StatusOK, w.Code)

		var received response.ListFlatsResponse
		ok := json.Unmarshal(w.Body.Bytes(), &received)
		require.Nil(t, ok, "Failed to unmarshal response body")

		expected := response.CreateListFlatsResponse(origin)

		assert.Equal(t, expected, received, "Expected %v, got %v", origin, received)
	})

	t.Run("GetHouseClient", func(t *testing.T) {
		pg := testhelper.SetupDBMock(t, setupGetHouseDBMock(t, 1, "client", origin))

		repo := repository.New(pg)
		c := cache.New()

		tok := token.New("")
		send := sender.New()

		houseService := house.New(repo)
		flatService := flat.New(repo, c)
		authService := auth.New(repo, tok)
		subService := subscribe.New(repo, send)
		hnd := handler.New(houseService, flatService, authService, subService)

		req, err := http.NewRequest("GET", "/house/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.WithValue(req.Context(), "role", "client")
		req = req.WithContext(ctx)

		w := httptest.NewRecorder()

		hnd.GetHouse(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d, got %d", http.StatusOK, w.Code)

		var received response.ListFlatsResponse
		ok := json.Unmarshal(w.Body.Bytes(), &received)
		require.Nil(t, ok, "Failed to unmarshal response body")

		expected := response.CreateListFlatsResponse(origin[2:3]) // only approved flats

		assert.Equal(t, expected, received, "Expected %v, got %v", origin, received)
	})

	t.Run("GetHouseClientFromCache", func(t *testing.T) {
		pg := &sqlx.DB{}

		repo := repository.New(pg)
		c := cache.New()

		err := c.PutHouse("1", origin[2:3])
		require.Nil(t, err)

		tok := token.New("")
		send := sender.New()

		houseService := house.New(repo)
		flatService := flat.New(repo, c)
		authService := auth.New(repo, tok)
		subService := subscribe.New(repo, send)
		hnd := handler.New(houseService, flatService, authService, subService)

		req, err := http.NewRequest("GET", "/house/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.WithValue(req.Context(), "role", "client")
		req = req.WithContext(ctx)

		w := httptest.NewRecorder()

		hnd.GetHouse(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d, got %d", http.StatusOK, w.Code)

		var received response.ListFlatsResponse
		ok := json.Unmarshal(w.Body.Bytes(), &received)
		require.Nil(t, ok, "Failed to unmarshal response body")

		expected := response.CreateListFlatsResponse(origin[2:3]) // only approved flats

		assert.Equal(t, expected, received, "Expected %v, got %v", origin, received)
	})

}

func TestCreateFlat(t *testing.T) {
	t.Run("ValidFlat", func(t *testing.T) {
		logger.MustInit("local")

		origin := domain.Flat{
			Id:      1,
			HouseId: 1,
			Price:   1000,
			Rooms:   3,
			Status:  "created",
		}
		pg := testhelper.SetupDBMock(t, setupCreateFlatDBMock(t, origin))

		repo := repository.New(pg)
		c := cache.New()

		tok := token.New("")
		send := sender.New()

		houseService := house.New(repo)
		flatService := flat.New(repo, c)
		authService := auth.New(repo, tok)
		subService := subscribe.New(repo, send)
		hnd := handler.New(houseService, flatService, authService, subService)

		body := []byte(`{"house_id":1,"price":1000,"rooms":3}`)
		req, err := http.NewRequest("POST", "/flats", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		hnd.CreateFlat(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d, got %d", http.StatusCreated, w.Code)

		var received response.Flat
		ok := json.Unmarshal(w.Body.Bytes(), &received)
		require.Nil(t, ok, "Failed to unmarshal response body")

		assert.EqualValues(t, origin, received, "Expected %v, got %v", origin, received)
	})
}

func setupGetHouseDBMock(tb testing.TB, id int, role string, flats []domain.Flat) testhelper.SetupDBMockOption {
	tb.Helper()

	return func(tb testing.TB, mock sqlmock.Sqlmock) {
		tb.Helper()

		rows := []string{
			"id",
			"house_id",
			"price",
			"rooms",
			"status",
		}

		allRows := sqlmock.NewRows(rows)
		approvedRows := sqlmock.NewRows(rows)

		for _, f := range flats {
			allRows.AddRow(
				f.Id,
				f.HouseId,
				f.Price,
				f.Rooms,
				f.Status,
			)
			if f.Status == "approved" {
				approvedRows.AddRow(
					f.Id,
					f.HouseId,
					f.Price,
					f.Rooms,
					f.Status,
				)
			}
		}

		switch role {
		case "moderator":
			mock.ExpectQuery("SELECT id, house_id, price, rooms, status FROM flat WHERE house_id = $1").
				WithArgs(id).
				WillReturnRows(allRows)
		case "client":
			mock.ExpectQuery("SELECT id, house_id, price, rooms, status FROM flat WHERE house_id = $1 AND status = 'approved'").
				WithArgs(id).
				WillReturnRows(approvedRows)
		}
	}
}

func setupCreateFlatDBMock(tb testing.TB, flat domain.Flat) testhelper.SetupDBMockOption {
	tb.Helper()

	return func(tb testing.TB, mock sqlmock.Sqlmock) {
		tb.Helper()

		rows := sqlmock.NewRows([]string{
			"id",
			"house_id",
			"price",
			"rooms",
			"status",
		})

		rows.AddRow(
			flat.Id,
			flat.HouseId,
			flat.Price,
			flat.Rooms,
			flat.Status,
		)

		mock.ExpectBegin()
		mock.ExpectQuery("INSERT INTO flat (house_id, price, rooms) VALUES ($1, $2, $3) RETURNING id, house_id, price, rooms, status").
			WithArgs(flat.HouseId, flat.Price, flat.Rooms).
			WillReturnRows(rows)
		mock.ExpectExec("UPDATE house SET updated_at = now() WHERE id = $1").
			WithArgs(flat.HouseId).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
	}
}
