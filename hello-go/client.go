package main

import (
	"log"

	"github.com/pynezz/hello-go/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func connectClient() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client!",
	}
	response, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling SayHello %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}
