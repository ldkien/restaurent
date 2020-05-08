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

func Register(w http.ResponseWriter, r *http.Request) {
	result := make(chan []byte)
	var registerRequest pb.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	go connectRegisterModule(registerRequest, result)
	w.Write(<-result)

}

func connectRegisterModule(registerRequest pb.RegisterRequest, result chan []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer close(result)
	defer cancel()
	response, err := client.Client.Register(ctx, &registerRequest)
	if err != nil {
		log.Logger.Error("could not greet: ", err)
		error := backendError.GetError(backendError.CONNECTOR_ERROR)
		response = &pb.RegisterResponse{
			Error: error,
		}
	}

	data, _ := entity.Marshaler.MarshalToString(response)

	result <- []byte(data)
}
