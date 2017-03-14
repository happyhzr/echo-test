package main

import (
	"github.com/insisthzr/echo-test/cookbook/twitter/conf"
	"github.com/insisthzr/echo-test/cookbook/twitter/db"
	"github.com/insisthzr/echo-test/cookbook/twitter/handlers"

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
		return c.String(200, "pong")
	})
	e.GET("/api/v1/db/ping", func(c echo.Context) error {
		err := db.CheckStatus()
		if err != nil {
			return err
		}
		return c.String(200, "pong")
	})

	e.POST("/api/v1/signup", handlers.Signup)
	e.POST("/api/v1/login", handlers.Login)

	v1 := e.Group("/api/v1", middleware.JWTWithConfig(conf.JWTConfig))

	v1.POST("/follow/:id", handlers.Follow)
	//v1.POST("/posts", handlers.CreatePost)
	//v1.POST("/feed", handlers.FetchPost)

	e.Logger.Fatal(e.Start(":1323"))
}
