package shared

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
		// Password: cfg.Redis.Pass
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
