package utils

import (
	"time"

	"github.com/insisthzr/echo-test/cookbook/twitter/conf"

	"github.com/dgrijalva/jwt-go"
)

func NewToken(id string) *jwt.Token {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token
}

func NewSignedString(token *jwt.Token) (string, error) {
	signedString, err := token.SignedString([]byte(conf.SIGNING_KEY))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func UserIDFromToken(token *jwt.Token) string {
	claims := token.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
