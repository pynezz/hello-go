package main

import (
	"fmt"
	"log"
	"net"

	//chat "github.com/pynezz/chat"
	chat "github.com/pynezz/hello-go/chat"
	hello "github.com/pynezz/hello-go/hello_world"
	basics "github.com/pynezz/hello-go/the_basics"

	"github.com/pynezz/hello-go"	"google.golang.org/grpc"
)


var age int
var name string = "Kevin"

func main() {
	// runPackageHello_Go()
	// printFromBasics(name, 24)
	message, err := hello.RandomGreeting(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	fmt.Println(("Starting gRPC server..."))
	startServer()
	//scraper()
}

func startServer() {
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

	fmt.Println("Server started")
}

func printFromBasics(name string, age int) {
	basics.PrintPerson(name, age)
}

func printString() {
	age = 24
	fmt.Printf("\nMy age: %d \n", age)
	added := addTwoInts(2, 2)
	fmt.Printf("2 + 2 = %v", added)
}

func addTwoInts(i int, j int) int {
	add := func(i int, j int) int {
		k := i + j
		return k
	}
	return add(i, j)
}

func runPackageHello_Go() {
	h := hello.Hello{
		Hello: "Hey",
	}
	h.Hello_Go()

	printString()
}

// Notes:
