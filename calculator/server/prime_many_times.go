package main

import (
	"fmt"
	"log"
	"math"

	pb "github.com/flvritski/gRPC/calculator/proto"
)

func (s *Server) PrimeManyTimes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeManyTimesServer) error {
	log.Printf("PrimeManyTimes function was invoked with: %v\n", in)

	number := in.PrimeNumber
	divisor := int32(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				PrimeResult: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	// for i := 0; i < 10; i++ {
	// 	if isPrime(in.PrimeNumber) == true {
	// 		stream.Send(&pb.PrimeResponse{
	// 			PrimeResult: 1,
	// 		})
	// 	}
	// }
	return nil
}

func isPrime(n int32) bool {
	if n <= 1 {
		fmt.Println("number must be greater than 2.")
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	sqrtN := int(math.Sqrt(float64(n)))

	var i int32
	for i = 5; i <= int32(sqrtN); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
