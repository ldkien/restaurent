package handler

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"net/http"
	"restaurant/backend-base/entity"
	backendError "restaurant/backend-base/error"
	log "restaurant/backend-base/logger"
	pb "restaurant/backend-entity/entities"
	"restaurant/backend-router/grpc/client"
	"time"
)

func parseObjBeforeConnect(w http.ResponseWriter, r *http.Request, request proto.Message, baseRequest *entity.BaseRequest) bool {
	err := json.NewDecoder(r.Body).Decode(baseRequest)
	if err != nil {
		log.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	byte, err := baseRequest.Data.MarshalJSON()
	s := string(byte)
	err = jsonpb.UnmarshalString(s, request)
	if err != nil {
		log.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	return true
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	result := make(chan []byte)
	var baseRequest entity.BaseRequest
	var orderRequest pb.OrderRequest
	if !parseObjBeforeConnect(w, r, &orderRequest, &baseRequest) {
		return
	}
	orderRequest.Common = baseRequest.Common
	go connectModule(TYPE(0), orderRequest, result)
	w.Write(<-result)

}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	result := make(chan []byte)
	var baseRequest entity.BaseRequest
	var orderRequest pb.OrderRequest
	if !parseObjBeforeConnect(w, r, &orderRequest, &baseRequest) {
		return
	}
	orderRequest.Common = baseRequest.Common
	go connectModule(TYPE(1), orderRequest, result)
	w.Write(<-result)

}

func connectModule(orderType TYPE, orderRequest pb.OrderRequest, result chan []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer close(result)
	defer cancel()
	var response *pb.BaseResponse
	var err error
	switch orderType {
	case CREATE_ORDER:
		response, err = client.ClientOrder.Order(ctx, &orderRequest)
	case UPDATE_ORDER:
		response, err = client.ClientOrder.UpdateOrder(ctx, &orderRequest)

	}
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

func connectCreateOrderModule(orderRequest pb.OrderRequest, result chan []byte) {
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

func connectUpdateOrderModule(orderRequest pb.OrderRequest, result chan []byte) {
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

type TYPE int32

const (
	CREATE_ORDER TYPE = 0
	UPDATE_ORDER TYPE = 1
)
