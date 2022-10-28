package domain

import "github.com/dgrijalva/jwt-go"

type Token string

type HeaderWithToken struct {
	Authorization string
}

type JwtCustomClaims struct {
	Uid      int    `json:"uid"`
	UserName string `json:"name"`
	jwt.StandardClaims
}
