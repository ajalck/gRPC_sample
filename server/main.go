package main

import (
	"log"
	"net"

	pb "github.com/ajalck/gRPC_sample/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)
type helloServer struct{
	pb.GreetServiceServer
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start :%v", err)
	}
}
