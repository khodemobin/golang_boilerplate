package redis

import (
	"github.com/go-redis/redis/v9"
	"github.com/khodemobin/golang_boilerplate/internal/config"
)

func New(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Database,
		PoolSize: cfg.Redis.PoolSize,
	})
	return client
}
