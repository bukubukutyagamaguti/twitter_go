package infrastructure

import (
	"api/server/domain"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenHandler struct {
	Method *jwt.SigningMethodHMAC
	Secret string
}

func NewTokenHandler() *TokenHandler {
	tokenHandler := &TokenHandler{
		Method: jwt.SigningMethodHS256,
		Secret: os.Getenv("SECRET_KEY"),
	}
	if tokenHandler.Secret == "" {
		panic("secret key not defined.")
	}
	return tokenHandler
}

func (handler *TokenHandler) Generate(uid int, username string) (string, error) {
	jwtClaims := domain.JwtCustomClaims{
		Uid:      uid,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(),
		},
	}
	// set header
	token := jwt.NewWithClaims(handler.Method, jwtClaims)

	// signature
	tokenString, err := token.SignedString([]byte(handler.Secret))
	fmt.Printf("%v %v", tokenString, err)
	if err != nil {
		return tokenString, errors.New("failer to generate a new token." + err.Error())
	}
	return tokenString, nil
}

func (handler *TokenHandler) VerityToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(handler.Secret), nil
	})

	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if !ok {
			return errors.New("couldn't handle this token:" + err.Error())
		}

		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return errors.New("not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return errors.New("token is either expired or not active yet")
		} else {
			return errors.New("Couldn't handle this token:" + err.Error())
		}
	}

	return nil
}
