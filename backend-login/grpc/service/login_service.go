package service

import (
	"github.com/gocql/gocql"
	"restaurant/backend-base/app"
	"restaurant/backend-base/database/cassandra"
	backendEntity "restaurant/backend-base/entity"
	backendError "restaurant/backend-base/error"
	pb "restaurant/backend-entity/entities"
	"restaurant/backend-login/module"
	"strings"
)

func Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	response := getUserData(in)

	return response, nil
}

func getUserData(request *pb.LoginRequest) *pb.LoginResponse {
	query := "SELECT * FROM user_data where username = ? and password = ? ALLOW FILTERING"
	iterator := cassandra.Session.
		Query(query, strings.ToLower(request.Username), request.Password).Consistency(gocql.One).Iter()
	m := map[string]interface{}{}
	var found = false
	claims := &backendEntity.Claims{}
	for iterator.MapScan(m) {
		found = true
		claims.Username = app.ConvertInterfaceToString(m["username"])
		claims.Group = app.ConvertInterfaceToString(m["group"])
	}
	var errorResponse *pb.Error
	var token string
	if !found {
		errorResponse = backendError.GetError(backendError.LOGIN_ERROR)
	} else {
		errorResponse = backendError.GetError(backendError.SUCCESS)
		token = module.CreateJwtToken(claims)
		if len(token) == 0 {
			errorResponse = backendError.GetError(backendError.SYSTEM_ERROR)
		}
	}

	return &pb.LoginResponse{
		Error: errorResponse,
		Token: token,
	}
}
