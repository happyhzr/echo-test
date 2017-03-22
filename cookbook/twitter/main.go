package main

import (
	"net/http"

	"github.com/insisthzr/echo-test/cookbook/twitter/conf"
	"github.com/insisthzr/echo-test/cookbook/twitter/controllers"
	"github.com/insisthzr/echo-test/cookbook/twitter/db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	err := db.InitDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/api/v1/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.GET("/api/v1/db/ping", func(c echo.Context) error {
		err := db.CheckStatus()
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "pong")
	})

	e.POST("/api/v1/signup", controllers.Signup)
	e.POST("/api/v1/login", controllers.Login)

	v1 := e.Group("/api/v1", middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf.SigningKey),
	}))

	v1.POST("/follow/:to", controllers.Follow)
	v1.POST("/posts", controllers.CreatePost)
	v1.GET("/posts", controllers.FetchPost)

	e.Logger.Fatal(e.Start(":1323"))
}
