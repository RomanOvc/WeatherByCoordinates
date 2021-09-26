package handlers

import (
	"WeatherByCoordinates/auth"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
)

//sign key
var mySigningKey = []byte("ROMAN")

type Token struct {
	Token string `json:"token"`
}

func (authR *AuthHandler) CheckLogin(u *auth.User) (*Token, error) {
	user, err := authR.AuthR.GetUser(u.Username)
	if err != nil {
		return nil, err
	}
	if user.Username != u.Username || user.Password != u.Password {
		fmt.Println("NOT CORRECT")
		err := "error" // TODO rename
		return nil, errors.Errorf(err) // TODO wrap
	}

	validToken, err := generateJWT(user)
	// fmt.Println(validToken)
	if err != nil {
		fmt.Println(err)
	}

	return &Token{validToken}, err

}

// генерация токена
func generateJWT(u *auth.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = u.User_id
	claims["username"] = u.Username
	claims["password"] = u.Password // TODO ???

	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		errors.Wrapf(err, "err token")
	}
	return tokenString, err
}