package main

import (
	"githubclitracker/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("error: %v", err)
	}

}
