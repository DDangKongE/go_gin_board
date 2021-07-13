package service

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"go_board/src/Models"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

func Sign(user *Models.User) jwt.Token {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("failed to generate private key: %s\n", err)
		return nil
	}

	var payload []byte
	{ // Create signed payload
		token := jwt.New()
		token.Set(`user_id`, user.USER_ID)
		token.Set(`user_email`, user.USER_EMAIL)
		token.Set(`user_nickname`, user.USER_NICKNAME)

		payload, err = jwt.Sign(token, jwa.RS256, privKey)
		if err != nil {
			fmt.Printf("failed to generate signed payload: %s\n", err)
			return nil
		}
		a_t, err2 := jwt.Parse(payload)
		fmt.Println("asf", a_t)
		fmt.Println("asf2", err2)

		return a_t
	}
}
