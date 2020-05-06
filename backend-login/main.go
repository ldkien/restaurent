package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"restaurant/backend-base/app"
	log "restaurant/backend-base/logger"
	pb "restaurant/backend-entity/entities"
	"restaurant/backend-login/grpc/service"
)

const (
	port = ":8001"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func main() {
	lis, err := net.Listen("tcp", app.GlobalConfig.LoginPort)
	if err != nil {
		log.Logger.Error("failed to listen: ", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Logger.Error("failed to serve: ", err)
	}
	log.Logger.Info("Start module login port: " + app.GlobalConfig.LoginPort)
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return service.Login(in)
}
