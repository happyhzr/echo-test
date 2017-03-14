package handlers

import (
	"time"

	"github.com/insisthzr/echo-test/cookbook/twitter/conf"
	"github.com/insisthzr/echo-test/cookbook/twitter/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func Signup(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		return err
	}

	if user.Email == "" || user.Password == "" {
		return c.String(401, "email or password is nil")
	}

	exist, err := models.UserExist(user.Email)
	if err != nil {
		return err
	}
	if exist {
		return c.String(400, "user exist")
	}

	err = user.AddUser()
	if err != nil {
		return err
	}

	return c.JSON(201, user)
}

func Login(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		return err
	}

	valid, err := user.ValidUser()
	if err != nil {
		return err
	}
	if !valid {
		return c.String(401, "email or password is incorrect")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	user.Token, err = token.SignedString([]byte(conf.SigningKey))
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

func Follow(c echo.Context) error {
	followerID := userIDFromToken(c)
	id := c.Param("id")

	err := models.AddFollower(bson.ObjectIdHex(id), followerID)
	if err != nil {
		return err
	}
	return c.NoContent(200)
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
