package handlers

import (
	"time"

	"github.com/insisthzr/echo-test/cookbook/twitter/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) Signup(c echo.Context) error {
	user := &models.User{ID: bson.NewObjectId()}
	err := c.Bind(user)
	if err != nil {
		return err
	}

	if user.Email == "" || user.Password == "" {
		return &echo.HTTPError{Code: 401, Message: "invalid email or password"}
	}

	db := h.DB.Clone()
	defer db.Close()
	err = db.DB("twitter").C("users").Insert(user)
	if err != nil {
		return err
	}

	return c.JSON(201, user)
}

func (h *Handler) Login(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		return err
	}

	db := h.DB.Clone()
	defer db.Close()
	err = db.DB("twitter").C("users").
		Find(bson.M{"email": user.Email, "password": user.Password}).
		One(user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{301, "invalid email or password"}
		}
		return err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	user.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	user.Password = ""
	return c.JSON(200, user)
}

func (h *Handler) Follow(c echo.Context) error {
	userID := userIDFromToken(c)
	id := c.Param("id")

	db := h.DB.Clone()
	defer db.Close()
	err := db.DB("twitter").C("users").
		UpdateId(bson.ObjectIdHex(id), bson.M{"$addToSet": bson.M{"followers": userID}})
	if err == mgo.ErrNotFound {
		return echo.ErrNotFound
	}
	return err
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
