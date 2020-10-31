package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/alileza/example/autogen/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewExampleServiceV1Client(conn)
	resp, err := client.Status(context.Background(), &pb.Empty{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("MESSAGE: %+v\n", resp)
}
