package main

import (
	"log"
	"net"
	"os"

	"grpc-example-with-go/internal/app"
	handler "grpc-example-with-go/internal/handler/grpc"
	gen "grpc-example-with-go/internal/handler/grpc/generated"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}

	svc := app.NewProductService()
	handler := handler.NewProductGrpcHandler(svc)
	server := grpc.NewServer()

	gen.RegisterProductHandlerServer(server, handler)

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}
}
