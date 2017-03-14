package conf

import "github.com/labstack/echo/middleware"

const (
	DBName = "twitter"

	SigningKey = "secret"
)

var (
	JWTConfig = middleware.JWTConfig{
		SigningKey: []byte(SigningKey),
	}
)
