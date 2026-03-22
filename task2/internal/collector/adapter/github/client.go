package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang-course/task2/internal/collector/domain"
)

type apiRepo struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stars       int       `json:"stargazers_count"`
	Forks       int       `json:"forks_count"`
	CreatedAt   time.Time `json:"created_at"`
}

type Client struct {
	HTTP *http.Client
}

func (c *Client) Fetch(ctx context.Context, owner, repo string) (*domain.Repository, error) {
	httpClient := c.HTTP
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return nil, domain.ErrRepositoryNotFound
	case http.StatusOK:
		break
	default:
		return nil, fmt.Errorf("github api: unexpected status %s", resp.Status)
	}

	var body apiRepo
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	return &domain.Repository{
		Name:        body.Name,
		Description: body.Description,
		Stars:       body.Stars,
		Forks:       body.Forks,
		CreatedAt:   body.CreatedAt,
	}, nil
}
