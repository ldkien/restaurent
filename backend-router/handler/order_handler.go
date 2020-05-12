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

func Order(w http.ResponseWriter, r *http.Request) {
	/*result := make(chan []byte)
	var baseRequest entity.BaseRequest
	err := json.NewDecoder(r.Body).Decode(&baseRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var orderRequest pb.OrderRequest
	err = ptypes.UnmarshalAny(baseRequest.Data, &orderRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	go connectOrderModule(orderRequest, result)
	w.Write(<-result)*/

}

func connectOrderModule(orderRequest pb.OrderRequest, result chan []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer close(result)
	defer cancel()
	response, err := client.ClientOrder.Order(ctx, &orderRequest)
	if err != nil {
		log.Logger.Error("could not greet: ", err)
		error := backendError.GetError(backendError.CONNECTOR_ERROR)
		response = &pb.BaseResponse{
			Error: error,
		}
	}

	data, _ := entity.Marshaler.MarshalToString(response)

	result <- []byte(data)
}
