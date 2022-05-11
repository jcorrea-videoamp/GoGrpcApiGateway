package main

import (
	"fmt"
	"log"
	"net"

	proto "github.com/jcorrea-videoamp/GoGrpcApiGateway/project/proto/v1"
	"github.com/jcorrea-videoamp/GoGrpcApiGateway/project/service"
	"google.golang.org/grpc"
)

func main() {
	srv, _ := service.NewDummyService()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterDummyServiceServer(server, srv)
	server.Serve(lis)
}
