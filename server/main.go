package main

import (
	"log"
	"net"

	pb "github.com/afthab/microservices/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8081"
)

type helloServer struct {
	pb.GreetServiceServer //This indicates that your helloServer struct will implement the methods defined in the GreetServiceServer interface.
}

func main() {
	lis, err := net.Listen("tcp", port) //used for low-level network operations, such as creating listeners and accepting connections,
	if err != nil {
		log.Fatalf("Failed to Start the Server %v", err)
	}
	grpcServer := grpc.NewServer() //starts a grpc server
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server Started at %v ", lis.Addr())
	err1 := grpcServer.Serve(lis) //This starts the server and begins listening for incoming gRPC requests.
	if err1 != nil {
		log.Fatalf("Failed to start : %v", err1)
	}
	//This line registers your implementation of the gRPC service interface (helloServer) with the gRPC server (grpcServer). It ensures that incoming gRPC requests for the GreetService will be handled by the methods implemented in the helloServer struct.

}
