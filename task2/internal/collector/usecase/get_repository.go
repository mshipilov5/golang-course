package usecase

import (
	"context"
	"strings"

	"golang-course/task2/internal/collector/domain"
)

type GetRepository struct {
	Fetcher domain.RepositoryFetcher
}

func (g *GetRepository) Run(ctx context.Context, owner, repo string) (*domain.Repository, error) {
	owner = strings.TrimSpace(owner)
	repo = strings.TrimSpace(repo)
	if owner == "" || repo == "" {
		return nil, domain.ErrInvalidInput
	}
	return g.Fetcher.Fetch(ctx, owner, repo)
}
