package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/flvritski/gRPC/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	//doSum(c)
	//doPrimeManyTimes(c)
	//doAvg(c)
	//doMax(c)
	//doSqrt(c, 10)
	doSqrt(c, -2)
}
