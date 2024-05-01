package main

import (
	"log"
	"net"

	"github.com/Animesh-roy100/chat-server-grpc/chatserver"
	"google.golang.org/grpc"
)

func main() {
	// Create a TCP listener at port 5000
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Error creating the server: %v", err)
	}

	log.Println("Server started at port :5000")

	// gRPC server instance
	grpcServer := grpc.NewServer()

	// register ChatService
	// create an instance of ChatServer and pass it along with gRPC instance to RegisterServiceServer method
	cs := &chatserver.ChatServer{}
	chatserver.RegisterServicesServer(grpcServer, cs)

	// gRPC listen and serve
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to start gPRC server: %v", err)
	}
}
