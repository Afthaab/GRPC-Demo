package main

import (
	"context"
	"log"
	"time"

	pb "github.com/afthab/microservices/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming Started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	for _, name := range names.Name {
		req := &pb.HelloRequest{
			Name: name,
		}
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the request with name : %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err1 := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err1 != nil {
		log.Fatalf("Error while closing the stream %v", err1)
	}
	log.Printf("%v", res.Name)
}
