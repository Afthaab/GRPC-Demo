package main

import (
	"log"

	pb "github.com/afthab/microservices/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8081"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials())) // grpc.Dial to create a connection to the gRPC server. The address passed to grpc.Dial is the server's network address and port localhost is concatenated with the port variable.

	//grpc.WithTransportCredentials is used to specify the transport credentials for the gRPC connection. In this case, insecure.NewCredentials() provides an insecure connection, typically used for testing or local development. In a production environment, you would use secure credentials such as TLS.

	if err != nil {
		log.Fatalf("Connection Failed : %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Name: []string{"Akhil", "Alice", "Bob"},
	}

	// callSayHello(client)
	// callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}
