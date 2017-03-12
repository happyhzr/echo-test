package main

import (
	"echo-test/cookbook/twitter/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handlers.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/v1/login" || c.Path() == "/v1/signup" {
				return true
			}
			return false
		},
	}))

	db, err := mgo.Dial("localhost")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Create indices
	err = db.Copy().DB("twitter").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	})
	if err != nil {
		e.Logger.Fatal(err)
	}

	h := &handlers.Handler{DB: db}

	v1 := e.Group("/v1")

	v1.POST("/signup", h.Signup)
	v1.POST("/login", h.Login)
	v1.POST("/follow/:id", h.Follow)
	v1.POST("/posts", h.CreatePost)
	v1.POST("/feed", h.FetchPost)

	e.Logger.Fatal(e.Start(":1323"))
}
