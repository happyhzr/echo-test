package controllers

import (
	"github.com/insisthzr/echo-test/cookbook/twitter/models"
	"github.com/insisthzr/echo-test/cookbook/twitter/utils"

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

	user.ID = bson.NewObjectId()
	err = user.AddUser()
	if err != nil {
		return err
	}

	user.Password = ""
	return c.JSON(201, user)
}

func Login(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		return err
	}
	if user.Email == "" || user.Password == "" {
		return c.String(400, "email or password is nil")
	}

	existUser, err := models.FindUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existUser == nil {
		return c.String(400, "user not exist")
	}
	if user.Password != existUser.Password {
		return c.String(400, "email or password is incorrect")
	}

	token := utils.NewToken(existUser.ID.String())
	signedString, err := utils.NewSignedString(token)
	if err != nil {
		return err
	}
	existUser.Token = signedString
	existUser.Password = "" // don't send password

	return c.JSON(200, existUser)
}

func Follow(c echo.Context) error {
	from := utils.UserIDFromToken(c.Get("user").(*jwt.Token))
	to := c.Param("to")

	err := models.AddFollower(bson.ObjectIdHex(to), from)
	if err != nil {
		return err
	}
	return c.JSON(200, map[string]string{
		"from": from,
		"to":   to,
	})
}
