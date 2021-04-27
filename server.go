package main

/*
	CREDIT: TutorialEdge [Beginners Guide to gRPC in Go!] - https://www.youtube.com/watch?v=BdzYdN_Zd9Q
*/

import (
	"log"
	"net"

	chat "github.com/pynezz/chat"
	"google.golang.org/grpc"
)

func StartServer() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.UnimplementedChatServiceServer{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
