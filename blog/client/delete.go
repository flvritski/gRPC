package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/flvritski/gRPC/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	fmt.Println("---deleteBlog was invoked---")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Println("Blog was deleted\n")
}
