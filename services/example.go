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
