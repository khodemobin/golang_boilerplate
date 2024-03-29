package server

import (
	"fmt"
	"github.com/khodemobin/golang_boilerplate/internal/app"
	"github.com/khodemobin/golang_boilerplate/internal/server/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	app *fiber.App
}

func New(isLocal bool) *Server {
	return &Server{
		app: fiber.New(fiber.Config{
			Prefork: !isLocal,
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				app.Log().Error(err)
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			},
		}),
	}
}

func (r *Server) Start(isLocal bool, port int) error {
	if isLocal {
		r.app.Use(fiberLogger.New())
	} else {
		r.app.Use(recover.New(), compress.New())
		r.app.Use(middleware.NewSentryFiber(middleware.Options{
			Repanic: true,
		}))
	}

	r.app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	r.routing()
	return r.app.Listen(fmt.Sprintf(":%d", port))
}

func (r *Server) Shutdown() error {
	return r.app.Shutdown()
}
