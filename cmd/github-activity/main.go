package main

import (
	"flag"
	"fmt"
	"github.com/letsmakecakes/github-activity/internal/api"
	"github.com/letsmakecakes/github-activity/internal/formatter"
	"log"
	"os"
)

func main() {
	// Parse command-line arguments
	flag.Parse()
	username := flag.Arg(0)

	// Ensure a username is provided
	if username == "" {
		printUsageAndExit()
	}

	// Create a new GitHub client
	client := api.NewGitHubClient()

	// Fetch user events from GitHub
	events, err := client.FetchUserEvents(username)
	if err != nil {
		log.Fatalf("Error fetching events: %v", err)
	}

	// Format and print the events
	output := formatter.FormatEvents(events)
	fmt.Print(output)
}

// printUsageAndExit prints the usage message and exits the program
func printUsageAndExit() {
	fmt.Println("Please provide a GitHub username")
	fmt.Println("Usage: github-activity <username>")
	os.Exit(1)
}
