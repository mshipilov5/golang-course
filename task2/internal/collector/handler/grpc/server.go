package grpchandler

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "golang-course/task2/api/proto/collectorv1"
	"golang-course/task2/internal/collector/domain"
	"golang-course/task2/internal/collector/usecase"
)

type Server struct {
	pb.UnimplementedCollectorServiceServer
	UC *usecase.GetRepository
}

func (s *Server) GetRepositoryInfo(ctx context.Context, req *pb.GetRepositoryInfoRequest) (*pb.GetRepositoryInfoResponse, error) {
	repo, err := s.UC.Run(ctx, req.GetOwner(), req.GetRepo())
	if err != nil {
		return nil, mapError(err)
	}
	return &pb.GetRepositoryInfoResponse{
		Name:        repo.Name,
		Description: repo.Description,
		Stars:       int32(repo.Stars),
		Forks:       int32(repo.Forks),
		CreatedAt:   repo.CreatedAt.Format(time.RFC3339),
	}, nil
}

func mapError(err error) error {
	switch {
	case errors.Is(err, domain.ErrRepositoryNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, domain.ErrInvalidInput):
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
