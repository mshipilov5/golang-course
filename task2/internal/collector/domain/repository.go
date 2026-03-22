package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrRepositoryNotFound = errors.New("repository not found")
	ErrInvalidInput       = errors.New("owner and repo are required")
)

type Repository struct {
	Name        string
	Description string
	Stars       int
	Forks       int
	CreatedAt   time.Time
}

type RepositoryFetcher interface {
	Fetch(ctx context.Context, owner, repo string) (*Repository, error)
}
