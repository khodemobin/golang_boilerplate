package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pio/provider/internal/domain"
	"github.com/khodemobin/pio/provider/pkg/logger"
)

type Sample struct {
	Logger  logger.Logger
	Service domain.SampleService
}

func (h *Sample) Sample(c *fiber.Ctx) error {
	return c.JSON("")
}
