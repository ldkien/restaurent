package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gocql/gocql"
	"restaurant/backend-base/app"
	"restaurant/backend-base/database/cassandra"
	backendEntity "restaurant/backend-base/entity"
	backendError "restaurant/backend-base/error"
	"restaurant/backend-base/logger"
	pb "restaurant/backend-entity/entities"
	"time"
)


func Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	response := getUserData(in)

	return response, nil
}

func getUserData(request *pb.LoginRequest) *pb.LoginResponse {
	query := "SELECT * FROM user_data where username = ? and password = ? ALLOW FILTERING"
	iterator := cassandra.Session.
		Query(query, request.Username, request.Password).Consistency(gocql.One).Iter()
	m := map[string]interface{}{}
	var found = false
	claims := &backendEntity.Claims{}
	for iterator.MapScan(m) {
		found = true
		claims.Username = m["username"].(string)
		claims.Group = m["group"].(string)
	}
	var errorResponse *pb.Error
	var token string
	if !found {
		errorResponse = backendError.GetError(backendError.LOGIN_ERROR)
	}else {
		errorResponse = backendError.GetError(backendError.SUCCESS)
		token = createJwtToken(claims)
		if len(token) == 0 {
			errorResponse = backendError.GetError(backendError.SYSTEM_ERROR)
		}
	}

	return &pb.LoginResponse{
		Error: errorResponse,
		Token: token,
	}
}

func createJwtToken(claims *backendEntity.Claims) string  {
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