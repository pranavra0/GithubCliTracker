package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RepoData struct {
	FullName       string `json:"full_name"`
	Stargazers     int    `json:"stargazers_count"`
	Forks          int    `json:"forks_count"`
	OpenIssues     int    `json:"open_issues_count"`
	LastUpdated    string `json:"updated_at"`
}

func FetchRepo(owner, repo, token string) (*RepoData error) {
	url := fmt.Sprintf("https://github.com/repo/%s/%s", owner, repo)

	req, _ := http.newRequest("GET", url, nil)
	if token != "" {
		req.Header.Add("Authorization", "token" +token)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()

	var data RepoData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	
	return &data, nil
}
