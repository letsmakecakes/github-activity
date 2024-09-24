# GitHub Activity CLI

A simple command-line tool to fetch and display a GitHub user's recent activity using the [GitHub API](https://docs.github.com/en/rest/reference/activity). This project helps you practice working with APIs, handling JSON data, and building a CLI in Go.

## Features

- Fetches and displays the recent activity of a GitHub user in the terminal.
- CLI accepts the GitHub username as an argument.
- Handles errors gracefully, including invalid usernames or API failures.
- (Optional) Caching mechanism to improve performance.

## Getting Started

### Prerequisites

- Go 1.18+
- GitHub API token (if needed for higher rate limits)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/<your-username>/github-activity-cli.git
   cd github-activity-cli
   ```

2. Install the necessary dependencies:

   ```bash
   go mod tidy
   ```

3. Build the CLI tool:

   ```bash
   go build -o github-activity ./cmd/github-activity
   ```

### Usage

1. Run the tool by passing a GitHub username as an argument:

   ```bash
   ./github-activity <username>
   ```

   Example:

   ```bash
   ./github-activity kamranahmedse
   ```

   This will display the recent public activity of the specified user, like:

   ```
   - Pushed 3 commits to kamranahmedse/developer-roadmap
   - Opened a new issue in kamranahmedse/developer-roadmap
   - Starred kamranahmedse/developer-roadmap
   ```

### Error Handling

The tool will gracefully handle common errors like:

- Invalid GitHub usernames
- Network issues or GitHub API failures
- Rate-limiting by the GitHub API

In case of errors, appropriate messages will be displayed in the terminal.

### Optional: Caching

You can enable caching to store the fetched activity data temporarily, reducing redundant API calls for the same user.

## Project Structure

```bash
github-activity-cli/
├── cmd/
│   └── github-activity/
│       └── main.go                # CLI entry point
├── internal/
│   ├── api/
│   │   └── github.go              # GitHub API requests
│   └── model/
│       └── event.go               # Data model for GitHub events
├── pkg/
│   └── cli/
│       └── cli.go                 # CLI utilities or reusable logic
├── go.mod                         # Go module dependencies
├── go.sum                         # Go module checksums
└── README.md                      # Project documentation
```

## API Reference

This tool uses the following GitHub API endpoint to fetch user activity:

```
GET https://api.github.com/users/<username>/events
```

For more details on the GitHub API, see the [official documentation](https://docs.github.com/en/rest/reference/activity).

## Future Enhancements

- **Filtering by event type**: Filter events such as push, issue creation, starring, etc.
- **Structured output**: Display activity in a more organized or structured format.
- **Additional API endpoints**: Fetch more user-related data, like repositories, followers, etc.
