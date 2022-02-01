package server

import (
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	handler "github.com/khodemobin/pio/provider/internal/server/handlers"
	"github.com/khodemobin/pio/provider/internal/service"
	"github.com/khodemobin/pio/provider/pkg/logger"
)

type Server struct {
	app    *fiber.App
	sample *handler.Sample
}

func New(service *service.Service, isLocal bool, logger logger.Logger) *Server {
	return &Server{
		app: fiber.New(fiber.Config{
			Prefork: !isLocal,
		}),
		sample: &handler.Sample{
			Logger:  logger,
			Service: service.Sample,
		},
	}
}

func (r *Server) Start(isLocal bool, port string) error {
	if isLocal {
		r.app.Use(fiberLogger.New())
		r.app.Use(recover.New())
	}

	r.routing()
	return r.app.Listen(":" + port)
}

func (r *Server) Shutdown() error {
	return r.app.Shutdown()
}

func (r *Server) routing() {
	r.app.Post("/sample", r.sample.Sample)
}
