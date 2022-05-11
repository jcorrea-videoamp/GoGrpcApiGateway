package main

import (
	"context"
	"log"
	"net/http"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	proto "github.com/jcorrea-videoamp/GoGrpcApiGateway/project/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Connect to the GRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	// Register grpc-gateway
	rmux := runtime.NewServeMux()
	client := proto.NewDummyServiceClient(conn)
	err = proto.RegisterDummyServiceHandlerClient(ctx, rmux, client)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("REST server ready...")
	err = http.ListenAndServe("localhost:8080", rmux)
	if err != nil {
		log.Fatal(err)
	}
}
