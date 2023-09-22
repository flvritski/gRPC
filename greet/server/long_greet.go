package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/flvritski/gRPC/greet/proto"
)

func (s *Server) GreetLong(stream pb.GreetService_GreetLongServer) error {
	log.Println("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
