package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"house-service/internal/domain"
	"house-service/internal/http/handler/mocks"
	"house-service/internal/http/model/response"
	"house-service/internal/logger"
	"house-service/pkg/utils/dbErrors"
)

func TestCreateFlat(t *testing.T) {
	log := logger.New("prod")

	t.Run("ValidRequest", func(t *testing.T) {
		flatService := &mocks.FlatService{}
		flatService.On("CreateFlat", context.Background(), 1, 1000, 3).
			Return(&domain.Flat{
				Id:      1,
				HouseId: 1,
				Price:   1000,
				Rooms:   3,
			}, nil)

		h := &Handler{
			flatService: flatService,
			log:         log,
		}

		body := []byte(`{"house_id":1,"price":1000,"rooms":3}`)
		req := httptest.NewRequest("POST", "/flat", bytes.NewReader(body))
		w := httptest.NewRecorder()

		h.CreateFlat(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d, got %d", http.StatusCreated, w.Code)

		expectedResponse := response.Flat{
			Id:      1,
			HouseId: 1,
			Price:   1000,
			Rooms:   3,
		}

		var received response.Flat
		ok := json.Unmarshal(w.Body.Bytes(), &received)
		require.Nil(t, ok, "Failed to unmarshal response body")

		assert.Equal(t, expectedResponse, received)

		assert.EqualValues(t, expectedResponse, received)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		flatService := &mocks.FlatService{}

		h := &Handler{
			flatService: flatService,
			log:         log,
		}

		req := httptest.NewRequest("POST", "/flat", strings.NewReader("invalid json"))
		w := httptest.NewRecorder()

		h.CreateFlat(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	})

	t.Run("FlatServiceError", func(t *testing.T) {
		flatService := &mocks.FlatService{}
		flatService.On("CreateFlat", context.Background(), 1, 1000, 3).
			Return(nil, errors.New(dbErrors.ErrFailedConnection))

		h := &Handler{
			flatService: flatService,
			log:         log,
		}

		body := []byte(`{"house_id":1,"price":1000,"rooms":3}`)
		req := httptest.NewRequest("POST", "/flat", bytes.NewReader(body))
		w := httptest.NewRecorder()

		h.CreateFlat(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code, "Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	})

	//t.Run("InvalidFlat", func(t *testing.T) {
	//	flatService := &mocks.FlatService{}
	//	flatService.On("CreateFlat", context.Background(), 1, 1000, 3).
	//		Return(nil, dbErrors.ErrAlreadyExists)
	//
	//	h := &Handler{
	//		flatService: flatService,
	//		log:         log,
	//	}
	//
	//	body := []byte(`{"house_id":1,"price":1000,"rooms":3}`)
	//	req := httptest.NewRequest("POST", "/flat", bytes.NewReader(body))
	//	w := httptest.NewRecorder()
	//
	//	h.CreateFlat(w, req)
	//
	//	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	//})
}

func TestGetHouse(t *testing.T) {
	log := logger.New("prod")

	t.Run("ValidRequest", func(t *testing.T) {
		flatService := &mocks.FlatService{}
		flatService.On("GetHouse", context.Background(), 1).
			Return([]domain.Flat{
				{
					Id:      1,
					HouseId: 123,
					Price:   100000,
					Rooms:   3,
					Status:  "approved",
				},
			}, nil)

		h := &Handler{
			flatService: flatService,
			log:         log,
		}

		req := httptest.NewRequest("GET", "/house/1", nil)
		w := httptest.NewRecorder()

		h.GetHouse(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "Expected status code %d, got %d", http.StatusOK, w.Code)

		expectedResponse := response.ListFlatsResponse{
			Flats: []response.Flat{
				{
					Id:      1,
					HouseId: 123,
					Price:   100000,
					Rooms:   3,
					Status:  "approved",
				},
			},
		}

		var received response.ListFlatsResponse
		ok := json.Unmarshal(w.Body.Bytes(), &received)
		require.Nil(t, ok, "Failed to unmarshal response body")

		assert.Equal(t, expectedResponse, received)
	})

	t.Run("HouseNotFound", func(t *testing.T) {
		flatService := &mocks.FlatService{}
		flatService.On("GetHouse", context.Background(), 1).
			Return(nil, errors.New(dbErrors.ErrNotFound))

		h := &Handler{
			flatService: flatService,
			log:         log,
		}

		req := httptest.NewRequest("GET", "/house/1", nil)
		w := httptest.NewRecorder()

		h.GetHouse(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code, "Expected status code %d, got %d", http.StatusNotFound, w.Code)

		expectedError := "entity not found"
		assert.Equal(t, expectedError, strings.TrimSpace(w.Body.String())) // TODO: ???
	})

	t.Run("ErrorGettingHouse", func(t *testing.T) {
		flatService := &mocks.FlatService{}
		flatService.On("GetHouse", context.Background(), 1).
			Return(nil, errors.New(dbErrors.ErrFailedConnection))

		h := &Handler{
			flatService: flatService,
			log:         log,
		}

		req := httptest.NewRequest("GET", "/house/1", nil)
		w := httptest.NewRecorder()

		h.GetHouse(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code, "Expected status code %d, got %d", http.StatusInternalServerError, w.Code)

		expectedError := response.Error{
			Message: "failed connection",
			//RequestId: "1",
			Code: 500,
		}

		var received response.Error
		ok := json.Unmarshal(w.Body.Bytes(), &received)
		require.Nil(t, ok, "Failed to unmarshal response body")

		assert.Equal(t, expectedError, received, "Expected error %v, got %v", expectedError, received)
	})
}
