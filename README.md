# GitHub Activity CLI

A simple command-line interface tool that fetches and displays GitHub user activity using the GitHub API.

## Features

- Fetch recent GitHub activity for any user
- Display activity in a clean, readable format
- Support for multiple event types:
   - Push events (with commit counts)
   - Issue events (opened/closed)
   - Watch events (starring repositories)
- Built with pure Go - no external dependencies
- Error handling for API failures and invalid usernames
- Clean and maintainable codebase structure

## Installation

### Prerequisites

- Go 1.21 or higher
- Git

### Building from Source

1. Clone the repository
```bash
git clone https://github.com/yourusername/github-activity.git
cd github-activity
```

2. Build the project
```bash
go build -o github-activity cmd/github-activity/main.go
```

3. (Optional) Add to PATH
```bash
# On Unix-like systems
sudo mv github-activity /usr/local/bin/
```

## Usage

```bash
github-activity <username>
```

Example:
```bash
$ github-activity kamranahmedse
- Pushed 3 commits to kamranahmedse/developer-roadmap
- Opened an issue in kamranahmedse/developer-roadmap
- Starred kamranahmedse/developer-roadmap
```

## Project Structure

```
github-activity/
├── cmd/
│   └── github-activity/
│       └── main.go           # Application entry point
├── internal/
│   ├── api/
│   │   └── github.go         # GitHub API client
│   ├── models/
│   │   └── event.go         # Data structures
│   └── formatter/
│       └── output.go        # Output formatting logic
├── go.mod
└── README.md
```

## Error Handling

The tool handles various error cases:

- Invalid GitHub usernames
- API rate limiting
- Network connectivity issues
- Malformed API responses

Example error output:
```bash
$ github-activity invalid@user
Error fetching events: GitHub API error: 404 Not Found - {"message": "Not Found"}
```

## Contributing

Contributions are welcome! Here are some ways you can contribute:

1. Report bugs
2. Suggest new features
3. Submit pull requests

### Development Setup

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Possible Enhancements

- [ ] Add authentication support for higher API rate limits
- [ ] Implement event filtering by type
- [ ] Add caching for better performance
- [ ] Support more event types
- [ ] Add pagination support
- [ ] Include colorized output
- [ ] Add detailed event information
- [ ] Support custom date ranges

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- GitHub API Documentation
- Go standard library
- The open-source community

## Support

If you encounter any problems or have suggestions, please open an issue in the GitHub repository.