package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/insisthzr/echo-test/cookbook/twitter/db"
	"github.com/insisthzr/echo-test/cookbook/twitter/models"
	"github.com/insisthzr/echo-test/cookbook/twitter/utils"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	post = models.Post{"", "", "someonefrom", "something message"}
)

func TestCreatePost(t *testing.T) {
	err := db.InitDB()
	if err != nil {
		t.Fatal(err)
	}
	e := echo.New()
	b, err := json.Marshal(post)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(echo.POST, "/posts", strings.NewReader(string(b)))
	if assert.NoError(t, err) {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		from := "58c75a9ff825dda245aa7aa4" // email=hello
		token := utils.NewToken(from)
		c.Set("user", token)
		if assert.NoError(t, CreatePost(c)) {
			t.Log(rec.Body.String())
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	}
}

func TestFetchPost(t *testing.T) {
	err := db.InitDB()
	if err != nil {

	}
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/posts", nil)
	if assert.NoError(t, err) {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		to := "58c790831d41c8f869f03013" // email=hello1
		token := utils.NewToken(to)
		c.Set("user", token)

		if assert.NoError(t, FetchPost(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	}
}
