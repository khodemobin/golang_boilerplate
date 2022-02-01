package service

import (
	"github.com/khodemobin/pio/provider/internal/domain"
	"github.com/khodemobin/pio/provider/internal/repository"
	"github.com/khodemobin/pio/provider/pkg/logger"
	"github.com/khodemobin/pio/provider/pkg/messager"
)

type Service struct {
	Sample domain.SampleService
}

func NewService(repo *repository.Repository, logger logger.Logger, msg messager.Messager) *Service {
	sa := NewSampleService(logger)

	return &Service{
		Sample: sa,
	}
}
