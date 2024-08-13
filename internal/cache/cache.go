package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
	"house-service/internal/model"
)

const (
	defaultExpiration = 10 * time.Minute
	cleanupInterval   = 15 * time.Minute
)

type Cache struct {
	pool *cache.Cache
}

func New() *Cache {
	return &Cache{
		pool: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (c *Cache) PutHouse(id string, house []model.Flat) error {
	err := c.pool.Add(id, house, defaultExpiration)
	if err != nil {
		// TODO: ?
		return err
	}

	return nil
}

func (c *Cache) GetHouse(id string) ([]model.Flat, bool) {
	items, ok := c.pool.Get(id)

	house, ok := items.([]model.Flat)
	if !ok {
		return nil, false
	}

	return house, ok
}
