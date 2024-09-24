package cli

import (
	"fmt"
	"log"

	"github.com/letsmakecakes/github-activity/internal/api"
	"github.com/letsmakecakes/github-activity/internal/model"
)

func HandleCommand(args []string) {
	username := args[1]
	URL := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	reqObj := api.NewRequest(URL)
	resp, err := reqObj.SendClientRequest()
	if err != nil {
		log.Panicf("error sending request: %v", err)
	}

	respObj := api.NewResponse()
	err = respObj.GetEvents(resp)
	if err != nil {
		log.Panicf("error getting events: %v", err)
	}

	printEvents(respObj.Event)
}

func printEvents(events []model.Event) {
	for _, event := range events {
		for _, commit := range event.Payload.Commits {
			msg := commit.Message
			if len(msg) > 0 {
				fmt.Printf("- %s\n", msg)
			}
		}
	}
}
