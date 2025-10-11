# Project Summary

## What is TGBot4SillyTavern?

TGBot4SillyTavern is a Telegram bot that lets you interact with your SillyTavern characters directly from Telegram, without needing to use the web browser.

## Why Use This Bot?

✅ **Mobile-Friendly**: Chat with your characters from your phone via Telegram  
✅ **Always Accessible**: No need to keep a browser tab open  
✅ **Button Interface**: Easy-to-use buttons instead of typing commands  
✅ **Full Features**: Access all SillyTavern features through Telegram  
✅ **Privacy**: Self-hosted solution - your data stays on your server  

## What Can You Do?

### Core Features

1. **Character Management**
   - Browse all your characters
   - Switch between characters with one tap
   - See character names in an organized list

2. **Chat Functionality**
   - Send messages and get AI responses
   - Regenerate responses if you want a different answer
   - Edit previous messages
   - Full formatting support (bold, italic, links)

3. **History Management**
   - Access all your saved chats
   - Navigate through chat history with pagination
   - Load any previous conversation instantly

4. **Settings Control**
   - Switch between completion presets
   - Configure AI behavior on the fly
   - All without leaving Telegram

## How It Works

```
You → Telegram → Bot → Headless Browser → SillyTavern
```

The bot uses a headless Chrome browser to interact with SillyTavern's web interface, acting as a bridge between you and your characters. This means:

- ✅ Works with any SillyTavern setup
- ✅ No API changes needed
- ✅ Gets all frontend features automatically
- ✅ Updates work seamlessly with SillyTavern

## Technical Details

### Technology Stack
- **Language**: Go (Golang) 1.21+
- **Browser Automation**: chromedp
- **Telegram API**: go-telegram-bot-api
- **Deployment**: Docker & Docker Compose

### System Requirements
- **For Docker**: 2GB RAM, 1GB disk space
- **For Local**: Go 1.21+, Chrome/Chromium browser
- **Network**: Access to your SillyTavern instance

### Security & Privacy
- Self-hosted: runs on your own server
- No external dependencies
- Your data never leaves your infrastructure
- Open source: audit the code yourself

## Project Structure

```
TGBot4SillyTavern/
├── config/              # Configuration management
├── internal/
│   ├── bot/            # Telegram bot logic
│   ├── browser/        # Browser automation
│   └── formatter/      # Text formatting
├── .github/workflows/  # CI/CD pipelines
├── main.go            # Entry point
├── Dockerfile         # Container definition
└── docker-compose.yml # Easy deployment
```

## Documentation

- **[README.md](README.md)** - Main documentation and technical details
- **[QUICKSTART.md](QUICKSTART.md)** - Beginner-friendly setup guide (5 minutes)
- **[USAGE.md](USAGE.md)** - How to use the bot features
- **[CONTRIBUTING.md](CONTRIBUTING.md)** - How to contribute to the project
- **[CHANGELOG.md](CHANGELOG.md)** - Version history

## Getting Started

### For Beginners
Follow the [Quick Start Guide](QUICKSTART.md) - takes 5 minutes!

### For Developers
1. Clone the repo
2. Run `./setup.sh` 
3. Run `docker-compose up -d`
4. Check logs with `docker-compose logs -f`

## Roadmap

### Current Version (v1.0)
- ✅ Character selection
- ✅ Chat functionality
- ✅ History management
- ✅ Preset switching
- ✅ Rich text formatting

### Future Enhancements (Optional)
- 🔄 Edit completion presets via bot
- 🔄 World book management
- 🔄 Character card editing
- 🔄 Streaming responses
- 🔄 Multi-user support
- 🔄 Voice message support

## Support & Community

- **Issues**: [GitHub Issues](https://github.com/CambriaDev/TGBot4SillyTavern/issues)
- **Discussions**: [GitHub Discussions](https://github.com/CambriaDev/TGBot4SillyTavern/discussions)
- **Contributions**: See [CONTRIBUTING.md](CONTRIBUTING.md)

## License

MIT License - See [LICENSE](LICENSE) for details

## Acknowledgments

- **SillyTavern Team** - For the amazing character chat platform
- **chromedp** - For Go Chrome DevTools Protocol
- **go-telegram-bot-api** - For Telegram Bot API implementation

---

**Ready to get started?** → [Quick Start Guide](QUICKSTART.md)  
**Need help?** → [Usage Guide](USAGE.md)  
**Want to contribute?** → [Contributing Guide](CONTRIBUTING.md)
