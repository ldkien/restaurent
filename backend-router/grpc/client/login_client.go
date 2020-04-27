package client

import (
	"google.golang.org/grpc"
	"restaurant/backend-base/app"
	log "restaurant/backend-base/logger"
	pb "restaurant/backend-entity/entities"
)

var Client pb.GreeterClient
var conn *grpc.ClientConn

func init() {
	conn, err := grpc.Dial(app.GlobalConfig.LoginClient, grpc.WithInsecure(), grpc.WithInsecure())
	if err != nil {
		log.Logger.Error("did not connect: ", err)
	}
	Client = pb.NewGreeterClient(conn)
}

func Close()  {
	conn.Close()
}