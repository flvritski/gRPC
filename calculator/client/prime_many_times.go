package main

import (
	"context"
	"io"
	"log"

	pb "github.com/flvritski/gRPC/calculator/proto"
)

func doPrimeManyTimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimeManyTimes was invoked")

	req := &pb.PrimeRequest{
		PrimeNumber: 1239039284,
	}

	stream, err := c.PrimeManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling PrimeManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		log.Printf("PrimeManyTimes: %d\n", msg.PrimeResult)
	}
}
