package main

import (
	"context"
	"fmt"
	"log"

	"github.com/itsksaurabh/udemy/grpc/unary/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	// create request
	req := greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Kumar",
			LastName:  "Saurabh",
		},
	}

	// call Greet service
	res, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	fmt.Println(res.Result)
}
