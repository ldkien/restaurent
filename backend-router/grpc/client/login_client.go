package client

import (
	"google.golang.org/grpc"
	"restaurant/backend-base/app"
	"restaurant/backend-base/entity"
	log "restaurant/backend-base/logger"
	pb "restaurant/backend-entity/entities"
)

var Client pb.GreeterClient
var conn *grpc.ClientConn

func init() {
	conn, err := grpc.Dial("dummy", grpc.WithInsecure(), grpc.WithBalancer(grpc.RoundRobin(entity.NewPseudoResolver(app.GlobalConfig.LoginClient))))
	if err != nil {
		log.Logger.Error("did not connect: ", err)
	}
	Client = pb.NewGreeterClient(conn)
}

func Close() {
	conn.Close()
}
