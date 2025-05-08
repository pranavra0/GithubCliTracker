package main

import (
	"GithubCliTracker/cmd"
	"fmt"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}

}
