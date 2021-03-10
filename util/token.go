package util

import (
	"learn-gin/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["userId"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(config.Key["secret"]))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(token string) (string, error) {
	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Key["secret"]), nil
	})
	if err != nil {
		return "", err
	}
	claims := tk.Claims.(jwt.MapClaims)
	return claims["userId"].(string), nil
}
