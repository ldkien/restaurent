package service

import (
	"github.com/gocql/gocql"
	"restaurant/backend-base/app"
	"restaurant/backend-base/database/cassandra"
	backendEntity "restaurant/backend-base/entity"
	backendError "restaurant/backend-base/error"
	"restaurant/backend-base/logger"
	pb "restaurant/backend-entity/entities"
	"restaurant/backend-login/module"
	"strings"
	"time"
)

func Register(in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	response := validateDataRegister(in)
	if response != nil {
		return response, nil
	}

	response = checkUsernameExist(in.Username)
	if response != nil {
		return response, nil
	}

	response = insertDB(in)
	if response != nil {
		return response, nil
	}

	claims := &backendEntity.Claims{}
	claims.Username = in.Username
	claims.Group = app.DEFAULT_GROUP_USER
	token := module.CreateJwtToken(claims)
	if len(token) == 0 {
		logger.Logger.Info("Register create token failed")
		return &pb.RegisterResponse{
			Error: backendError.GetError(backendError.SYSTEM_ERROR),
			Token: "",
		}, nil
	}
	return &pb.RegisterResponse{
		Error: backendError.GetError(backendError.SUCCESS),
		Token: token,
	}, nil
}

func validateDataRegister(in *pb.RegisterRequest) *pb.RegisterResponse {
	in.Username = strings.TrimSpace(in.Username)
	in.Password = strings.TrimSpace(in.Password)
	in.RepeatPass = strings.TrimSpace(in.RepeatPass)
	in.FullName = strings.TrimSpace(in.FullName)
	if in == nil || len(in.Username) <= 0 || len(in.Password) <= 0 || len(in.RepeatPass) <= 0 || in.Password != in.RepeatPass {
		return &pb.RegisterResponse{
			Error: backendError.GetError(backendError.INVALID_PARAMS),
			Token: "",
		}
	}

	if !app.IsValidUsername(in.Username) {
		return &pb.RegisterResponse{
			Error: backendError.GetError(backendError.INVALID_USERNAME),
			Token: "",
		}
	}

	if in.FullName != "" && (len(in.FullName) > 50 || !app.IsLetter(in.FullName)) {
		return &pb.RegisterResponse{
			Error: backendError.GetError(backendError.INVALID_PARAMS),
			Token: "",
		}
	}

	in.Username = strings.ToLower(in.Username)
	return nil
}

func checkUsernameExist(username string) *pb.RegisterResponse {
	query := "SELECT username FROM user_data where username = ?"
	iterator := cassandra.Session.
		Query(query, username).Consistency(gocql.One).Iter()
	if iterator.NumRows() > 0 {
		return &pb.RegisterResponse{
			Error: backendError.GetError(backendError.EXIST_USERNAME),
			Token: "",
		}
	}
	return nil
}

func insertDB(in *pb.RegisterRequest) *pb.RegisterResponse {
	query := "INSERT INTO user_data (username, password, group, fullname, sex, createdate) VALUES (?, ?, ?, ?, ?, ?)"
	if err := cassandra.Session.Query(query,
		in.Username, in.Password, app.DEFAULT_GROUP_USER, in.FullName, pb.SEX_value[in.Sex.String()], time.Now()).Exec(); err != nil {
		logger.Logger.Error(err)
		return &pb.RegisterResponse{
			Error: backendError.GetError(backendError.DB_ERROR),
			Token: "",
		}
	}

	return nil
}
