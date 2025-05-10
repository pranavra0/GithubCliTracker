package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Repo struct {
	Owner string `yaml:"owner"`
	Name  string `yaml:"name"`
}

type Config struct {
	Repos []Repo `yaml:"repos"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	if len(cfg.Repos) == 0 {
		return nil, fmt.Errorf("no repos found in config")
	}

	return &cfg, nil
}
