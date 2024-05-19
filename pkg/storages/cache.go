package storage

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	USD_RATE_KEY = "rate:usd:uah"
)

type Cache struct {
	client *redis.Client
}

func NewCache(addr string) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &Cache{client: rdb}
}

func (r *Cache) SetRate(value interface{}) error {
	ctx := context.Background()
	expiration := 10 * time.Minute

	err := r.client.Set(ctx, USD_RATE_KEY, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("could not set value in Redis: %v", err)
	}
	return nil
}

func (r *Cache) GetRate() (float64, error) {
	ctx := context.Background()

	v, err := r.client.Get(ctx, USD_RATE_KEY).Result()
	if err != nil {
		return 0, err
	}

	f, err := strconv.ParseFloat(v, 64)

	if err != nil {
		return 0, err
	}

	return f, nil
}
