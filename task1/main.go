package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

type RepoInfo struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stars       int       `json:"stargazers_count"`
	Forks       int       `json:"forks_count"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	owner := flag.String("owner", "", "Owner of the GitHub repository")
	repo := flag.String("repo", "", "Name of the GitHub repository")

	flag.Parse()

	if *owner == "" || *repo == "" {
		fmt.Println("Usage: go run main.go -owner=<owner> -repo=<repo>")
		return
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", *owner, *repo)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		fmt.Println("Repository Not Found")
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error making request:", resp.Status)
		return
	}

	var info RepoInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		fmt.Println("Error parsing JSON:", err)
	}

	fmt.Println("Repository:", info.Name)
	fmt.Println("Description:", info.Description)
	fmt.Println("Stars:", info.Stars)
	fmt.Println("Forks:", info.Forks)
	fmt.Println("Created at:", info.CreatedAt)
}
