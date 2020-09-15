package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"restaurant/backend-base/app"
	log "restaurant/backend-base/logger"
	pb "restaurant/backend-entity/entities"
	"restaurant/backend-order/grpc/service"
)

type server struct {
	pb.UnimplementedOrderServiceServer
}

func main() {
	lis, err := net.Listen("tcp", app.GlobalConfig.OrderPort)
	if err != nil {
		log.Logger.Error("failed to listen: ", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &server{})
	log.Logger.Info("Start module order port: " + app.GlobalConfig.OrderPort)
	if err := s.Serve(lis); err != nil {
		log.Logger.Error("failed to serve: ", err)
	}
}

func (s *server) Order(ctx context.Context, in *pb.OrderRequest) (*pb.BaseResponse, error) {
	return service.Order(in)
}
