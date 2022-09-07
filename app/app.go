package app

import (
	redisDriver "github.com/go-redis/redis/v8"

	"github.com/khodemobin/golang_boilerplate/internal/config"
	"github.com/khodemobin/golang_boilerplate/pkg/cache"
	"github.com/khodemobin/golang_boilerplate/pkg/helper"
	"github.com/khodemobin/golang_boilerplate/pkg/logger"
	"github.com/khodemobin/golang_boilerplate/pkg/logger/sentry"
	"github.com/khodemobin/golang_boilerplate/pkg/logger/syslog"
	"github.com/khodemobin/golang_boilerplate/pkg/logger/zap"
	"github.com/khodemobin/golang_boilerplate/pkg/mysql"
	"github.com/khodemobin/golang_boilerplate/pkg/redis"

	"gorm.io/gorm"
)

type AppContainer struct {
	Cache  cache.Cache
	DB     *gorm.DB
	Redis  *redisDriver.Client
	Log    logger.Logger
	Config *config.Config
}

var Container *AppContainer = nil

func New() {
	config := config.New()

	var logger logger.Logger
	if helper.IsLocal() {
		logger = zap.New()
	} else if config.App.Env == "test" {
		logger = syslog.New()
	} else {
		logger = sentry.New(Container.Config)
	}

	db := mysql.New(config, logger).DB
	rc := redis.New(config, logger)
	cache := cache.New(rc, logger)

	Container = &AppContainer{
		Config: config,
		Log:    logger,
		DB:     db,
		Cache:  cache,
	}
}

func App() *AppContainer {
	return Container
}

func Cache() cache.Cache {
	return Container.Cache
}

func DB() *gorm.DB {
	return Container.DB
}

func Redis() *redisDriver.Client {
	return Container.Redis
}

func Log() logger.Logger {
	return Container.Log
}

func Config() *config.Config {
	return Container.Config
}
