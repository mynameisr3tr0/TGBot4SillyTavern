# TGBot4SillyTavern

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Build Status](https://github.com/CambriaDev/TGBot4SillyTavern/workflows/Build%20and%20Test/badge.svg)](https://github.com/CambriaDev/TGBot4SillyTavern/actions)
[![Docker](https://img.shields.io/badge/docker-supported-2496ED?style=flat&logo=docker)](https://hub.docker.com/)

A Telegram Bot that integrates with [SillyTavern](https://github.com/SillyTavern/SillyTavern), allowing users to interact with their characters through Telegram instead of the web interface.

## 🎯 Features

### Core Functionality
- **Character Selection**: Browse and select characters through inline keyboard buttons
- **Chat History Management**: View and load previous chat sessions with pagination support
- **Message Operations**:
  - Send messages and receive AI responses
  - Regenerate the last AI response
  - Edit previous messages
- **Completion Presets**: Switch between different completion presets on-the-fly
- **Rich Text Support**: HTML formatting from SillyTavern is converted to Telegram-compatible format (bold, italic, links)

### User Interface
- Button-based interaction for ease of use
- Main menu with quick access to all features
- Inline action buttons on each message (Regenerate, Edit)
- Paginated lists for chat history

## 🛠️ Technical Stack

- **Language**: Go (Golang)
- **Browser Automation**: chromedp (Headless Chrome)
- **Telegram API**: go-telegram-bot-api
- **Deployment**: Docker support included

## 📋 Prerequisites

- Go 1.21 or higher (for local development)
- Docker and Docker Compose (for containerized deployment)
- A running instance of SillyTavern
- A Telegram Bot Token (obtain from [@BotFather](https://t.me/botfather))

**New to this?** Check out the [Quick Start Guide for Beginners](QUICKSTART.md) 👈

## 🚀 Quick Start

### Method 1: Docker (Recommended)

1. Clone the repository:
```bash
git clone https://github.com/CambriaDev/TGBot4SillyTavern.git
cd TGBot4SillyTavern
```

2. Run the setup script (interactive):
```bash
./setup.sh
```

Or create `.env` manually based on `.env.example`:
```bash
cp .env.example .env
# Edit .env with your configuration
```

3. Edit `.env` and set your configuration:
```env
TELEGRAM_BOT_TOKEN=your_telegram_bot_token_here
SILLYTAVERN_URL=http://host.docker.internal:8000
HEADLESS_MODE=true
DEBUG=false
```

4. Build and run with Docker Compose:
```bash
docker-compose up -d
```

5. Check logs:
```bash
docker-compose logs -f
```

### Method 2: Local Development

1. Clone the repository:
```bash
git clone https://github.com/CambriaDev/TGBot4SillyTavern.git
cd TGBot4SillyTavern
```

2. Install dependencies:
```bash
go mod download
```

3. Run the setup script or set environment variables:
```bash
./setup.sh
# OR manually:
export TELEGRAM_BOT_TOKEN=your_telegram_bot_token_here
export SILLYTAVERN_URL=http://localhost:8000
export HEADLESS_MODE=true
```

4. Run the bot:
```bash
go run main.go
```

## 📝 Configuration

All configuration is done through environment variables:

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `TELEGRAM_BOT_TOKEN` | Your Telegram bot token from BotFather | - | Yes |
| `SILLYTAVERN_URL` | URL to your SillyTavern instance | `http://localhost:8000` | No |
| `HEADLESS_MODE` | Run browser in headless mode | `true` | No |
| `CHROMIUM_PATH` | Custom path to Chromium binary | (auto-detected) | No |
| `DEBUG` | Enable debug logging | `false` | No |

## 🎮 Usage

1. Start a chat with your bot on Telegram
2. Send `/start` or `/menu` to open the main menu
3. Use the buttons to:
   - **👤 Select Character**: Choose which character to chat with
   - **💬 Chat History**: Load previous chat sessions
   - **⚙️ Completion Presets**: Switch AI generation settings

4. Chat normally by sending text messages
5. Use inline buttons on responses to:
   - 🔄 **Regenerate**: Get a different response
   - ✏️ **Edit**: Modify the message

### Example Bot Interface

```
Welcome to SillyTavern Bot!

Choose an option:
┌────────────────────┐
│ 👤 Select Character│
├────────────────────┤
│ 💬 Chat History    │
├────────────────────┤
│ ⚙️ Completion Presets│
└────────────────────┘
```

After selecting a character, chat naturally:
```
You: Hello!

Character: *waves enthusiastically* 
Hi there! How are you doing today?

[🔄 Regenerate] [✏️ Edit]
```

## 🏗️ Project Structure

```
TGBot4SillyTavern/
├── config/              # Configuration management
│   └── config.go
├── internal/
│   ├── bot/            # Telegram bot logic
│   │   └── bot.go
│   ├── browser/        # Headless browser automation
│   │   └── browser.go
│   └── formatter/      # HTML to Telegram format converter
│       └── html.go
├── main.go             # Application entry point
├── Dockerfile          # Docker build configuration
├── docker-compose.yml  # Docker Compose configuration
├── go.mod              # Go module definition
├── go.sum              # Go dependencies checksums
├── .env.example        # Example environment variables
├── .gitignore          # Git ignore rules
└── README.md           # This file
```

## 🔧 How It Works

```
┌─────────────┐        ┌──────────────┐       ┌─────────────────┐
│   Telegram  │        │   TGBot4ST   │       │  SillyTavern    │
│     User    │◄──────►│  (Go + CDP)  │◄─────►│  (Web UI)       │
└─────────────┘        └──────────────┘       └─────────────────┘
                              │
                              ▼
                       ┌──────────────┐
                       │   Headless   │
                       │   Chromium   │
                       └──────────────┘
```

The bot uses a headless Chrome browser to interact with SillyTavern's frontend:

1. **Browser Automation**: chromedp drives a headless Chrome instance that loads the SillyTavern web interface
2. **DOM Manipulation**: The bot interacts with SillyTavern by clicking buttons, filling forms, and reading content from the page
3. **Message Relay**: User messages from Telegram are sent to SillyTavern, and responses are forwarded back to Telegram
4. **Format Conversion**: HTML responses from SillyTavern are converted to Telegram's HTML format

This approach allows the bot to leverage all of SillyTavern's frontend features without needing direct API access.

## 🐛 Troubleshooting

### Bot doesn't start
- Verify your `TELEGRAM_BOT_TOKEN` is correct
- Check that SillyTavern is running and accessible at the configured URL

### Browser automation fails
- Ensure Chrome/Chromium is installed (handled automatically in Docker)
- Try setting `HEADLESS_MODE=false` for debugging
- Check that SillyTavern's UI elements haven't changed (selectors may need updating)

### Messages not formatting correctly
- The HTML formatter supports basic formatting (bold, italic, links)
- Complex formatting may need additional conversion rules

### Connection issues in Docker
- Use `host.docker.internal` instead of `localhost` when SillyTavern runs on the host
- Ensure proper network configuration in docker-compose.yml

## 🔮 Future Enhancements (Optional Features)

- Edit completion presets through the bot
- Edit world books
- Edit character cards
- Streaming responses support
- Multi-user support with session management

## 📄 License

This project is open source. Please check the repository for license details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## 🙏 Acknowledgments

- [SillyTavern](https://github.com/SillyTavern/SillyTavern) - The amazing character chat interface
- [chromedp](https://github.com/chromedp/chromedp) - Chrome DevTools Protocol for Go
- [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) - Telegram Bot API for Go
