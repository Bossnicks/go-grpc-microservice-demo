package main

import (
	"log"
	"net"

	handler "github.com/Bossnicks/go-grpc-microservice-demo/services/orders/handler/orders"
	"github.com/Bossnicks/go-grpc-microservice-demo/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	log.Panicln("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)

}