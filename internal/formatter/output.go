package formatter

import (
	"fmt"
	"github.com/letsmakecakes/github-activity/internal/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

// FormatEvents formats a slice of events into a single string.
func FormatEvents(events []models.Event) string {
	var output strings.Builder

	for _, event := range events {
		if line := formatEvent(event); line != "" {
			output.WriteString(line + "\n")
		}
	}

	return output.String()
}

// formatEvent formats a single event based on its type.
func formatEvent(event models.Event) string {
	switch event.Type {
	case "PushEvent":
		return formatPushEvent(event)
	case "IssuesEvent":
		return formatIssuesEvent(event)
	case "WatchEvent":
		return formatWatchEvent(event)
	default:
		return ""
	}
}

// formatPushEvent formats a PushEvent.
func formatPushEvent(event models.Event) string {
	commitCount := len(event.Payload.Commits)
	return fmt.Sprintf(" - Pushed %d commits to %s", commitCount, event.Repo.Name)
}

// formatIssueEvent formats an IssuesEvent.
func formatIssuesEvent(event models.Event) string {
	c := cases.Title(language.English)
	action := c.String(event.Payload.Action)
	return fmt.Sprintf(" - %s an issue in %s", action, event.Repo.Name)
}

// formatWatchEvent formats a WatchEvent.
func formatWatchEvent(event models.Event) string {
	return fmt.Sprintf(" - Starred %s", event.Repo.Name)
}
