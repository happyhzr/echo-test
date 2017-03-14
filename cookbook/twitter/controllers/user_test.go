package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/insisthzr/echo-test/cookbook/twitter/db"

	"github.com/insisthzr/echo-test/cookbook/twitter/utils"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	err := db.InitDB()
	if !assert.NoError(t, err) {
		return
	}

	var userIn = map[string]string{
		"email":    "hello" + strconv.FormatInt(time.Now().Unix(), 10),
		"password": "world",
	}
	b, _ := json.Marshal(userIn)

	e := echo.New()
	req, err := http.NewRequest("POST", "/api/v1/signup", strings.NewReader(string(b)))
	if !assert.NoError(t, err) {
		return
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if !assert.NoError(t, Signup(c)) {
		return
	}

	assert.Equal(t, 201, rec.Code)
	t.Log("Signup:", rec.Body.String())
}

func TestLogin(t *testing.T) {
	err := db.InitDB()
	if !assert.NoError(t, err) {
		return
	}

	user := map[string]string{
		"email":    "hello",
		"password": "world",
	}
	b, _ := json.Marshal(user)

	e := echo.New()
	req, err := http.NewRequest("POST", "/api/v1/login", strings.NewReader(string(b)))
	if !assert.NoError(t, err) {
		return
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if !assert.NoError(t, Login(c)) {
		return
	}

	assert.Equal(t, 200, rec.Code)
	t.Log("Login:", rec.Body.String())
}

func TestFollow(t *testing.T) {
	err := db.InitDB()
	if !assert.NoError(t, err) {
		return
	}

	from := "58c75a9ff825dda245aa7aa4" // email=hello
	to := "58c790831d41c8f869f03013"   // email=hello1

	e := echo.New()
	req, err := http.NewRequest("POST", "/api/v1/follow/:to", nil)
	if !assert.NoError(t, err) {
		return
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := utils.NewToken(from)
	c.Set("user", token)
	c.SetParamNames("to")
	c.SetParamValues(to)

	if !assert.NoError(t, Follow(c)) {
		return
	}

	assert.Equal(t, 200, rec.Code)
	t.Log("follow:", rec.Body.String())
}
