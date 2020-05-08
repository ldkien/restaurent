package module

import (
	"github.com/dgrijalva/jwt-go"
	"restaurant/backend-base/app"
	backendEntity "restaurant/backend-base/entity"
	"restaurant/backend-base/logger"
	"time"
)

func CreateJwtToken(claims *backendEntity.Claims) string {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(app.JWT_KEY)
	if err != nil {
		logger.Logger.Error(err)
		tokenString = ""
	}
	return tokenString
}
