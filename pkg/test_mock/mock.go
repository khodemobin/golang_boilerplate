package test_mock

import (
	"os"
	"testing"

	"github.com/khodemobin/golang_boilerplate/app"
	"github.com/khodemobin/golang_boilerplate/pkg/cache"
	"github.com/khodemobin/golang_boilerplate/pkg/logger/syslog"
)

func NewMock(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	app.New()
	logger := syslog.New()
	redis := cache.NewTest(t, logger)
	app.Container.Cache = redis
}
