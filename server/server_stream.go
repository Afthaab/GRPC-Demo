package main

import (
	"log"
	"time"

	pb "github.com/afthab/microservices/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with the name %v", req.Name)
	for _, name := range req.Name {
		res := &pb.HelloResponse{
			Name: "hello" + name,
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil

}
