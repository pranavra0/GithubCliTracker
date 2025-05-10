package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"githubclitracker/config"
	"githubclitracker/github"
	"githubclitracker/tui"
)

var rootCmd = &cobra.Command{
	Use:   "githubclitracker",
	Short: "Track Github repo stats from terminal",
	Run: func(cmd *cobra.Command, args []string) {
		repos, err := config.LoadRepoList("repos.json")
		if err != nil {
			log.Fatalf("Failed to load repos: %v", err)
		}

		var repoData []*github.RepoData
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			log.Fatal("GITHUB_TOKEN not set in environment")
		}

		for _, repo := range repos {
			r, err := github.FetchRepo(repo, token)
			if err != nil {
				fmt.Printf("Error fetching repo %s: %v\n", repo, err)
				continue
			}
			repoData = append(repoData, r)
		}

		tui.RenderDashboard(repoData)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
