package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GithubCliTracker",
	Short: "Track Github repo stats from terminal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GithubCliTracker initialized. Add command logic here")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
