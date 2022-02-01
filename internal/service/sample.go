package service

import (
	"context"

	"github.com/khodemobin/pio/provider/internal/domain"
	"github.com/khodemobin/pio/provider/pkg/logger"
)

type sample struct {
	// repo   *repository.Repository
	logger logger.Logger
}

func NewSampleService(logger logger.Logger) domain.SampleService {
	return &sample{
		logger: logger,
	}
}

func (f *sample) Sample(ctx context.Context) (string, error) {
	return "", nil
}
