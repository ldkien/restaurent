package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"restaurant/backend-base/entity"
	backendError "restaurant/backend-base/error"
	log "restaurant/backend-base/logger"
	pb "restaurant/backend-entity/entities"
	"restaurant/backend-router/grpc/client"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	result := make(chan []byte)
	var loginRequest pb.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	go connectLoginModule(loginRequest, result)
	w.Write(<-result)

}

func connectLoginModule(loginRequest pb.LoginRequest, result chan []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer close(result)
	defer cancel()
	response, err := client.Client.Login(ctx, &loginRequest)
	if err != nil {
		log.Logger.Error("could not greet: ", err)
		error := backendError.GetError(backendError.CONNECTOR_ERROR)
		response = &pb.LoginResponse{
			Error: error,
		}
	}

	data, _ := entity.Marshaler.MarshalToString(response)

	result <- []byte(data)
}
