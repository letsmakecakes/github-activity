package main

import (
	"log"
	"os"

	"github.com/letsmakecakes/github-activity/pkg/cli"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Panic("expected a username")
	}

	cli.HandleCommand(args)
}
