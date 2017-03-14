package controllers

import (
	"strconv"

	"github.com/insisthzr/echo-test/cookbook/twitter/models"
	"github.com/insisthzr/echo-test/cookbook/twitter/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func CreatePost(c echo.Context) error {
	userID := utils.UserIDFromToken(c.Get("user").(*jwt.Token))
	post := new(models.Post)
	err := c.Bind(post)
	if err != nil {
		return err
	}

	post.ID = bson.NewObjectId()
	post.From = userID
	if post.To == "" || post.Message == "" {
		return &echo.HTTPError{Code: 400, Message: "invalid to or message"}
	}

	err = post.AddPost()
	if err != nil {
		return err
	}
	return c.JSON(201, post)
}

func FetchPost(c echo.Context) error {
	to := utils.UserIDFromToken(c.Get("user").(*jwt.Token))
	var err error

	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")
	page := 0
	limit := 100
	if len(pageStr) > 0 {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return c.String(400, err.Error())
		}
	}
	if len(limitStr) > 0 {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return c.String(400, err.Error())
		}
	}

	if page < 0 {
		page = 0
	}
	if limit < 0 {
		limit = 100
	}

	posts, err := models.FindPosts(to, page, limit)
	if err != nil {
		return err
	}

	return c.JSON(200, posts)
}
