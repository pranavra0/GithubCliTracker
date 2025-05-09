package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"githubclitracker/github"
	"githubclitracker/tui"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "githubclitracker",
	Short: "Track Github repo stats from terminal",
	Run: func(cmd *cobra.Command, args []string) {
		repos := []string{"torvalds/linux", "golang/go", "kubernetes/kubernetes"} // filler add other stuff later
		var repoData []*github.RepoData

		for _, repo := range repos {
			r, err := github.FetchRepo(repo, os.Getenv("GITHUB_TOKEN"))
			if err != nil {
				fmt.Println("Error fetching repo:", err)
				continue
			}
			repoData = append(repoData, r)
		}

		tui.RenderDashboard(repoData)
	}, // end of run
}

func Execute() error {
	return rootCmd.Execute()
}
