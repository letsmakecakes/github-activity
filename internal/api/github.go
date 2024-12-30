package api

import (
	"encoding/json"
	"fmt"
	"github.com/letsmakecakes/github-activity/internal/models"
	"io"
	"log"
	"net/http"
	"time"
)

// GitHubClient is a wrapper for making requests to the GitHub API.
type GitHubClient struct {
	client  *http.Client
	baseURL string
}

// NewGitHubClient creates a new instance of GitHubClient with default settings.
func NewGitHubClient() *GitHubClient {
	return &GitHubClient{client: &http.Client{
		Timeout: 10 * time.Second,
	},
		baseURL: "https://api.github.com",
	}
}

// FetchUserEvents retrieves the GitHub events for a given username.
func (c *GitHubClient) FetchUserEvents(username string) ([]models.Event, error) {
	url := fmt.Sprintf("%s/users/%s/events", c.baseURL, username)

	req, err := c.newRequest("GET", url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("error closing response body: %v", err)
		}
	}(resp.Body)

	if err := checkResponse(resp); err != nil {
		return nil, err
	}

	var events []models.Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return events, nil
}

// newRequest creates a new HTTP request with common headers set.
func (c *GitHubClient) newRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "GitHub-Activity-CLI")

	return req, nil
}

// checkResponse checks the HTTP response for errors.
func checkResponse(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("GitHub API error: %s - %s", resp.Status, string(body))
	}
	return nil
}
