package handlers

/*
import (
	"strconv"

	"github.com/insisthzr/echo-test/cookbook/twitter/models"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func CreatePost(c echo.Context) error {
	user := &models.User{
		ID: bson.ObjectIdHex(userIDFromToken(c)),
	}
	post := &models.Post{
		ID:   bson.NewObjectId(),
		From: user.ID.Hex(),
	}
	err := c.Bind(post)
	if err != nil {
		return err
	}

	if post.To == "" || post.Message == "" {
		return &echo.HTTPError{Code: 400, Message: "invalid to or message"}
	}

	db := h.DB.Clone()
	defer db.Close()

	err = db.DB("twitter").C("users").FindId(user.ID).One(user)
	if err != nil {
		return err
	}

	err = db.DB("twitter").C("posts").Insert(post)
	if err != nil {
		return err
	}
	return c.JSON(201, post)
}

func  FetchPost(c echo.Context) error {
	userID := userIDFromToken(c)
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	posts := make([]*models.Post, 0)
	db := h.DB.Clone()
	defer db.Close()

	err := db.DB("twitter").C("posts").
		Find(bson.M{"to": userID}).
		Skip((page - 1) * limit).
		Limit(limit).
		All(&posts)
	if err != nil {
		return err
	}

	return c.JSON(200, posts)

}
*/
