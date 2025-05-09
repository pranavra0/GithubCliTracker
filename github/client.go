package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RepoData struct {
	FullName    string `json:"full_name"`
	Stargazers  int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	OpenIssues  int    `json:"open_issues_count"`
	LastUpdated string `json:"updated_at"`
}

func FetchRepo(repo string, token string) (*RepoData, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s", repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "token "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch repo data: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var repoData RepoData
	if err := json.Unmarshal(body, &repoData); err != nil {
		return nil, err
	}

	return &repoData, nil
}

