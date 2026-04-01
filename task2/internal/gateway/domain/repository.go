package domain

import "time"

type RepositoryInfo struct {
	Name        string
	Description string
	Stars       int
	Forks       int
	CreatedAt   time.Time
}
