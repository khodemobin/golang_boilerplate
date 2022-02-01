package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/khodemobin/pio/provider/internal/cache"
	"github.com/khodemobin/pio/provider/internal/config"
	"github.com/khodemobin/pio/provider/internal/repository"
	"github.com/khodemobin/pio/provider/internal/server"
	"github.com/khodemobin/pio/provider/internal/service"
	"github.com/khodemobin/pio/provider/pkg/helper"
	"github.com/khodemobin/pio/provider/pkg/logger"
	"github.com/khodemobin/pio/provider/pkg/logger/sentry"
	"github.com/khodemobin/pio/provider/pkg/logger/zap"
	"github.com/khodemobin/pio/provider/pkg/messager/rabbit"
	"github.com/khodemobin/pio/provider/pkg/mysql"
	"github.com/khodemobin/pio/provider/pkg/redis"
)

func Execute() {
	// init main components
	config := config.New()

	var logger logger.Logger
	if helper.IsLocal(config) {
		logger = zap.New()
	} else {
		logger = sentry.New(config)
	}

	msg := rabbit.New(config, logger)
	db := mysql.New(config, logger)
	redis := redis.New(config, logger)

	cache := cache.New(redis, logger)

	defer db.Close()
	defer cache.Close()

	repository := repository.NewRepository(db.DB, cache)
	service := service.NewService(repository, logger, msg)

	// start server
	restServer := server.New(service, helper.IsLocal(config), logger)
	go func() {
		if err := restServer.Start(helper.IsLocal(config), config.App.Port); err != nil {
			msg := fmt.Sprintf("error happen while serving: %v", err)
			logger.Error(errors.New(msg))
			log.Println(msg)
		}
	}()

	// wait for close signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("Received an interrupt, closing connections...")

	if err := restServer.Shutdown(); err != nil {
		fmt.Println("Rest server doesn't shutdown in 10 seconds")
	}
}
