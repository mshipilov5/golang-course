package grpcadapter

import (
	"context"
	"time"

	"google.golang.org/grpc"

	pb "golang-course/task2/api/proto/collectorv1"
	"golang-course/task2/internal/gateway/domain"
)

type CollectorClient struct {
	client pb.CollectorServiceClient
}

func NewCollectorClient(conn *grpc.ClientConn) *CollectorClient {
	return &CollectorClient{client: pb.NewCollectorServiceClient(conn)}
}

func (c *CollectorClient) GetRepository(ctx context.Context, owner, repo string) (domain.RepositoryInfo, error) {
	resp, err := c.client.GetRepositoryInfo(ctx, &pb.GetRepositoryInfoRequest{
		Owner: owner,
		Repo:  repo,
	})
	if err != nil {
		return domain.RepositoryInfo{}, err
	}
	created, err := time.Parse(time.RFC3339, resp.GetCreatedAt())
	if err != nil {
		return domain.RepositoryInfo{}, err
	}
	return domain.RepositoryInfo{
		Name:        resp.GetName(),
		Description: resp.GetDescription(),
		Stars:       int(resp.GetStars()),
		Forks:       int(resp.GetForks()),
		CreatedAt:   created,
	}, nil
}
