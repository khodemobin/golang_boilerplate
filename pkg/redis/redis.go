package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/khodemobin/golang_boilerplate/internal/config"
	"github.com/khodemobin/golang_boilerplate/pkg/logger"
)

func New(cfg *config.Config, logger logger.Logger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Database,
		PoolSize: cfg.Redis.PoolSize,
	})

	return client
}
