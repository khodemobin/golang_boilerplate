package middleware

import (
	"net/http"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/golang_boilerplate/app"
	"github.com/khodemobin/golang_boilerplate/pkg/encrypt"
	"github.com/khodemobin/golang_boilerplate/pkg/helper"
)

var bearerRegexp = regexp.MustCompile(`^(?:B|b)earer (\S+$)`)

func JWTChecker(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusForbidden).JSON(helper.DefaultResponse(nil, "This endpoint requires a Bearer token", 0))
	}

	matches := bearerRegexp.FindStringSubmatch(authHeader)
	if len(matches) != 2 {
		return c.Status(http.StatusForbidden).JSON(helper.DefaultResponse(nil, "This endpoint requires a Bearer token", 0))
	}

	id, err := encrypt.ParseJWTClaims(matches[1])
	if err != nil {
		app.Log().Fatal(err)
	}
	c.Locals("user_id", id)

	return c.Next()
}
