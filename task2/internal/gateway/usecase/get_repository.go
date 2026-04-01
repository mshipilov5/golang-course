package usecase

import (
	"context"
	"golang-course/task2/internal/gateway/domain"
)

type RepositoryReader interface {
	GetRepository(ctx context.Context, owner, repo string) (domain.RepositoryInfo, error)
}

type GetRepository struct {
	Reader RepositoryReader
}

func (g *GetRepository) Run(ctx context.Context, owner, repo string) (domain.RepositoryInfo, error) {
	return g.Reader.GetRepository(ctx, owner, repo)
}
