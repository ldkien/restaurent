package entity

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Group string `json:"group"`
	jwt.StandardClaims
}
