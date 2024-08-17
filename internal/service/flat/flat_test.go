package flat

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"house-service/internal/domain"
	"house-service/internal/service/flat/mocks"
)

func TestCreateFlat(t *testing.T) {
	t.Run("ValidFlat", func(t *testing.T) {
		flatRepo := &mocks.FlatRepo{}
		flatService := &Service{repo: flatRepo}

		flatRepo.On("CreateFlat", context.Background(), 1, 1000, 3).
			Return(&domain.Flat{
				Id:      1,
				HouseId: 1,
				Price:   1000,
				Rooms:   3,
			}, nil)

		flat, err := flatService.CreateFlat(context.Background(), 1, 1000, 3)

		assert.NoError(t, err)
		assert.Equal(t, &domain.Flat{
			Id:      1,
			HouseId: 1,
			Price:   1000,
			Rooms:   3,
		}, flat)

		flatRepo.AssertCalled(t, "CreateFlat", context.Background(), 1, 1000, 3)
	})

	t.Run("InvalidFlat", func(t *testing.T) {
		flatRepo := &mocks.FlatRepo{}
		flatService := &Service{repo: flatRepo}

		flatRepo.On("CreateFlat", context.Background(), 1, 0, 3).
			Return(nil, errors.New("invalid flat"))

		flat, err := flatService.CreateFlat(context.Background(), 1, 0, 3)

		assert.Error(t, err)
		assert.Nil(t, flat)

		flatRepo.AssertCalled(t, "CreateFlat", context.Background(), 1, 0, 3)
	})
}

func TestGetHouse(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		flatRepo := &mocks.FlatRepo{}
		cache := &mocks.Cache{}
		flatService := &Service{repo: flatRepo, cache: cache}

		flatRepo.On("GetHouse", context.Background(), 1).
			Return([]domain.Flat{
				{
					Id:      1,
					HouseId: 1,
					Price:   1000,
					Rooms:   3,
					Status:  "approved",
				},
			}, nil)

		cache.On("GetHouse", "1").
			Return([]domain.Flat{
				{
					Id:      1,
					HouseId: 1,
					Price:   1000,
					Rooms:   3,
					Status:  "approved",
				},
			}, true)

		flats, err := flatService.GetHouse(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, []domain.Flat{
			{
				Id:      1,
				HouseId: 1,
				Price:   1000,
				Rooms:   3,
				Status:  "approved",
			},
		}, flats)

		//flatRepo.AssertCalled(t, "GetHouse", context.Background(), 1)
		//cache.AssertCalled(t, "GetHouse", "1")
	})

	t.Run("InvalidInput", func(t *testing.T) {
		flatRepo := &mocks.FlatRepo{}
		cache := &mocks.Cache{}
		flatService := &Service{repo: flatRepo, cache: cache}

		flatRepo.On("GetHouse", context.Background(), 0).
			Return(nil, errors.New("invalid input"))

		cache.On("GetHouse", "0").
			Return(nil, false)

		flats, err := flatService.GetHouse(context.Background(), 0)

		assert.Error(t, err)
		assert.Nil(t, flats)

		flatRepo.AssertCalled(t, "GetHouse", context.Background(), 0)
		cache.AssertCalled(t, "GetHouse", "0")
	})
}
