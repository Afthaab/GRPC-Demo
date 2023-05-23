package main

import (
	"io"
	"log"

	pb "github.com/afthab/microservices/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with the name: %v", req.Name)
		res := &pb.HelloResponse{
			Name: "Hello" + req.Name,
		}
		err1 := stream.Send(res)
		if err1 != nil {
			return err
		}
	}
}
