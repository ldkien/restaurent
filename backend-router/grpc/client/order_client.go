package client

import (
	"google.golang.org/grpc"
	"restaurant/backend-base/app"
	"restaurant/backend-base/entity"
	"restaurant/backend-base/logger"
	pb "restaurant/backend-entity/entities"
)

var ClientOrder pb.OrderServiceClient
var connOrder *grpc.ClientConn

func init() {
	connOrder, err := grpc.Dial("dummy", grpc.WithInsecure(), grpc.WithBalancer(grpc.RoundRobin(entity.NewPseudoResolver(app.GlobalConfig.OrderClient))))
	if err != nil {
		logger.Logger.Error("did not connect order service: ", err)
	}
	ClientOrder = pb.NewOrderServiceClient(connOrder)
}

func CloseOderClient() {
	if connOrder != nil {
		connOrder.Close()
	}
}
