package main

import (
	"context"
	"log"

	pb "github.com/flvritski/gRPC/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := pb.Blog{
		Id:       id,
		AuthorId: "Not Stefan",
		Title:    "New Title",
		Content:  "This is a new content version of the blog!",
	}

	_, err := c.UpdateBlog(context.Background(), &newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}
