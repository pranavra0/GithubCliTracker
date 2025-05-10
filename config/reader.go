package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadRepoList(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read repo config: %w", err)
	}

	var repos []string
	if err := json.Unmarshal(data, &repos); err != nil {
		return nil, fmt.Errorf("invalid repo config format: %w", err)
	}

	if len(repos) == 0 {
		return nil, fmt.Errorf("no repositories found in config")
	}

	return repos, nil
}
