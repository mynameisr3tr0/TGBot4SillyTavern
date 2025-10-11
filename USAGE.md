# Usage Guide

## Getting Your Telegram Bot Token

1. Open Telegram and search for [@BotFather](https://t.me/botfather)
2. Send `/newbot` command
3. Follow the prompts to name your bot
4. Copy the HTTP API token provided

## Setting Up SillyTavern

Make sure you have SillyTavern running and accessible:

1. Start SillyTavern on your machine
2. Note the URL (typically `http://localhost:8000`)
3. Ensure it's accessible from where the bot will run

## Running the Bot

### Option 1: Quick Start with Docker

```bash
# Create .env file
cat > .env << EOF
TELEGRAM_BOT_TOKEN=your_token_here
SILLYTAVERN_URL=http://host.docker.internal:8000
HEADLESS_MODE=true
DEBUG=false
EOF

# Run with docker-compose
docker-compose up -d

# View logs
docker-compose logs -f
```

### Option 2: Local Development

```bash
# Set environment variables
export TELEGRAM_BOT_TOKEN="your_token_here"
export SILLYTAVERN_URL="http://localhost:8000"

# Run the bot
go run main.go
```

## Bot Commands

- `/start` - Initialize the bot and show welcome message
- `/menu` - Display the main menu with all options
- `/characters` - Quick access to character selection
- `/history` - Quick access to chat history
- `/presets` - Quick access to completion presets

## Interaction Flow

### 1. Start Chatting

1. Open your bot on Telegram
2. Send `/start`
3. Select a character from the menu
4. Start chatting!

### 2. Character Management

```
/menu → 👤 Select Character → Choose from list
```

The bot will display all available characters from your SillyTavern instance.

### 3. Chat History

```
/menu → 💬 Chat History → Browse pages → Select chat
```

- Chat history is paginated (5 items per page)
- Use Previous/Next buttons to navigate
- Click on any chat to load it

### 4. Message Operations

Every AI response includes action buttons:

- **🔄 Regenerate**: Get a different response for the same prompt
- **✏️ Edit**: Edit the previous message (send new text after clicking)

### 5. Completion Presets

```
/menu → ⚙️ Completion Presets → Select preset
```

Switch between different AI generation settings on the fly.

## Tips & Tricks

### Rich Text Formatting

The bot supports:
- **Bold text** (from `<b>` or `<strong>`)
- *Italic text* (from `<i>` or `<em>`)
- [Links](url) (from `<a href="">`)
- Line breaks and paragraphs

### Docker on Same Machine as SillyTavern

If running in Docker on the same machine as SillyTavern:
```env
SILLYTAVERN_URL=http://host.docker.internal:8000
```

### Docker on Different Machine

If SillyTavern is on a different machine:
```env
SILLYTAVERN_URL=http://192.168.1.100:8000
```

### Debugging

Enable debug mode to see detailed logs:
```env
DEBUG=true
HEADLESS_MODE=false  # See browser window
```

## Troubleshooting

### "Failed to initialize browser"

**Solution**: Make sure chromium/chrome is installed
```bash
# On Ubuntu/Debian
sudo apt-get install chromium-browser

# Or use Docker which includes it
docker-compose up
```

### "Failed to navigate to SillyTavern"

**Solution**: Check SillyTavern is running and accessible
```bash
# Test from bot machine
curl http://localhost:8000
```

### "Failed to load characters"

**Possible causes**:
1. SillyTavern UI has been updated (selectors may need updating)
2. No characters exist in SillyTavern
3. Page hasn't fully loaded yet

**Solution**: Create a character in SillyTavern first, or increase wait times in browser code

### Messages not sending

**Check**:
1. Bot token is correct
2. You've started the bot with `/start`
3. A character is selected
4. SillyTavern is responsive

## Advanced Configuration

### Custom Chromium Path

If chromium is in a non-standard location:
```env
CHROMIUM_PATH=/usr/bin/chromium-browser
```

### Network Configuration

For complex network setups, modify `docker-compose.yml`:
```yaml
services:
  tgbot:
    network_mode: "host"  # Use host networking
    # ... rest of config
```

## Security Notes

- Never commit your `.env` file with real tokens
- Keep your bot token secret
- Consider using secrets management for production
- Run in a secure network environment

## Getting Help

If you encounter issues:

1. Check logs: `docker-compose logs -f`
2. Enable debug mode: `DEBUG=true`
3. Verify SillyTavern is accessible
4. Check GitHub issues for similar problems
5. Open a new issue with logs and configuration (remove sensitive data)
