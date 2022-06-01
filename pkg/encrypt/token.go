package encrypt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/khodemobin/golang_boilerplate/app"
	"github.com/khodemobin/golang_boilerplate/internal/model"
)

func GenerateAccessToken(user *model.User) (string, error) {
	secret := app.Config().Jwt.JwtSecret
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(app.Config().Jwt.JwtTTL)).Unix(),
		Subject:   fmt.Sprint(user.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseJWTClaims(bearer string) (string, error) {
	secret := app.Config().Jwt.JwtSecret

	p := jwt.Parser{ValidMethods: []string{jwt.SigningMethodHS256.Name}}
	c, err := p.ParseWithClaims(bearer, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	claims := c.Claims.(*jwt.StandardClaims)

	return claims.Subject, err
}
