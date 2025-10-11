# Contributing to TGBot4SillyTavern

Thank you for your interest in contributing to TGBot4SillyTavern! This document provides guidelines for contributing to the project.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/TGBot4SillyTavern.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Test your changes thoroughly
6. Commit and push to your fork
7. Open a Pull Request

## Development Setup

### Prerequisites
- Go 1.21 or higher
- Chrome/Chromium browser
- SillyTavern instance for testing
- Telegram Bot token

### Local Setup
```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/TGBot4SillyTavern.git
cd TGBot4SillyTavern

# Install dependencies
go mod download

# Copy environment template
cp .env.example .env

# Edit .env with your configuration
nano .env

# Run the bot
go run main.go
```

## Code Style

### Go Formatting
- Use `gofmt` to format your code
- Run `go vet` to check for common issues
- Use meaningful variable and function names
- Add comments for exported functions and complex logic

```bash
# Format code
go fmt ./...

# Check for issues
go vet ./...
```

### Project Structure
```
TGBot4SillyTavern/
├── config/          # Configuration management
├── internal/
│   ├── bot/        # Telegram bot logic
│   ├── browser/    # Browser automation
│   └── formatter/  # Text formatting utilities
└── main.go         # Application entry point
```

## Adding New Features

### Browser Automation
When adding new SillyTavern interactions:

1. Add methods to `internal/browser/browser.go`
2. Use proper error handling
3. Include adequate wait times for DOM elements
4. Test with both headless and headed modes

Example:
```go
func (b *Browser) NewFeature() error {
    err := chromedp.Run(b.ctx,
        chromedp.WaitVisible(`#element`, chromedp.ByID),
        chromedp.Click(`#element`, chromedp.ByID),
        chromedp.Sleep(500*time.Millisecond),
    )
    if err != nil {
        return fmt.Errorf("failed to execute new feature: %w", err)
    }
    return nil
}
```

### Bot Commands
When adding new Telegram bot features:

1. Add command handlers to `internal/bot/bot.go`
2. Create appropriate inline keyboards
3. Handle callback queries
4. Add user-friendly error messages

Example:
```go
case "new_feature":
    b.handleNewFeature(chatID)
```

### HTML Formatting
When adding new format conversions:

1. Update `internal/formatter/html.go`
2. Handle edge cases
3. Test with various HTML inputs
4. Preserve Telegram format limitations

## Testing

### Manual Testing
1. Test with a real SillyTavern instance
2. Verify all button interactions work
3. Test error scenarios
4. Check formatting with various message types

### Building
```bash
# Build the binary
make build

# Run with Docker
make docker-build
make docker-run
make docker-logs
```

## Pull Request Guidelines

### Before Submitting
- [ ] Code is formatted with `gofmt`
- [ ] All functions are properly documented
- [ ] Changes are tested manually
- [ ] No unnecessary dependencies added
- [ ] README updated if needed
- [ ] USAGE.md updated for new features

### PR Description
Include:
- What the PR does
- Why the change is needed
- How to test the changes
- Screenshots for UI changes
- Related issues (if any)

### Commit Messages
Use clear, descriptive commit messages:
```
Add character avatar display in selection menu

- Fetch avatar URLs from SillyTavern
- Display avatars in Telegram inline buttons
- Handle missing avatar gracefully
```

## Code Review Process

1. Maintainers will review your PR
2. Address any feedback or requested changes
3. Once approved, your PR will be merged
4. Your contribution will be acknowledged

## Areas for Contribution

### High Priority
- [ ] Improve error handling and user feedback
- [ ] Add retry logic for failed browser operations
- [ ] Optimize browser performance
- [ ] Add session persistence
- [ ] Multi-user support

### Medium Priority
- [ ] Edit completion presets through bot
- [ ] World book management
- [ ] Character card editing
- [ ] Export chat history

### Low Priority
- [ ] Streaming responses
- [ ] Voice message support
- [ ] Image generation support
- [ ] Advanced formatting options

## Reporting Bugs

### Before Reporting
1. Check existing issues
2. Verify the bug exists in the latest version
3. Collect relevant information

### Bug Report Template
```markdown
**Describe the bug**
A clear description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
1. Go to '...'
2. Click on '...'
3. See error

**Expected behavior**
What you expected to happen.

**Screenshots**
If applicable, add screenshots.

**Environment:**
- OS: [e.g., Ubuntu 22.04]
- Go version: [e.g., 1.21]
- SillyTavern version:
- Bot version:

**Logs**
Relevant log output (remove sensitive information)
```

## Feature Requests

We welcome feature requests! Please:
1. Check if the feature already exists or is planned
2. Describe the feature and its use case
3. Explain why it would be valuable
4. Provide examples if possible

## Questions?

- Open an issue for questions about the codebase
- Check existing documentation first
- Be specific and provide context

## License

By contributing, you agree that your contributions will be licensed under the same license as the project.

## Recognition

Contributors will be:
- Listed in the project README
- Mentioned in release notes for significant contributions
- Credited in commit history

Thank you for contributing! 🎉
