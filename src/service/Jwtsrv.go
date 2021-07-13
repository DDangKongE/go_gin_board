package service

import (
	"fmt"
	"go_board/src/Models"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func JwtService(user *Models.User) (string, error) {
	var err error

	os.Setenv("ACCESS_SECRET", "dkanrjskTma")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user"] = user
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	fmt.Println("ads", token)
	return token, nil
}
