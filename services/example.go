package services

import (
	"context"

	pb "github.com/alileza/example/autogen/pb"
)

type ExampleService struct{}

func (s *ExampleService) Status(ctx context.Context, _ *pb.Empty) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{
		Message: "Hello",
		Ok:      true,
	}, nil
}
func (s *ExampleService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &pb.HelloResponse{
		World: "world",
	}, nil
}
