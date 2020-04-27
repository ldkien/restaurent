package handler

import (
	"context"
	"net/http"
	"restaurant/backend-base/entity"
	backendError "restaurant/backend-base/error"
	log "restaurant/backend-base/logger"
	pb "restaurant/backend-entity/entities"
	"restaurant/backend-router/grpc/client"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.Client.Login(ctx, &pb.LoginRequest{Username: "kienmit", Password: "63ce66a3f0fd389fb9124826d6cdff29"})
	if err != nil {
		log.Logger.Error("could not greet: ", err)
		error := backendError.GetError(backendError.CONNECTOR_ERROR)
		response = &pb.LoginResponse{
			Error: error,
		}
	}

	data,_ := entity.Marshaler.MarshalToString(response)
	dataResponse := []byte(data)
	w.Write(dataResponse)
}
