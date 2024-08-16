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
	c.pool.DeleteExpired()

	items, ok := c.pool.Get(id)

	house, ok := items.([]model.Flat)
	if !ok {
		return nil, false
	}

	c.update(id, items)

	return house, ok
}

func (c *Cache) Delete(id string) {
	c.pool.Delete(id)
}

func (c *Cache) update(id string, items any) {
	err := c.pool.Replace(id, items, defaultExpiration)
	if err != nil {
		_ = c.pool.Add(id, items, defaultExpiration)
	}
}
