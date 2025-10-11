# Quick Start Guide for Beginners

This guide will help you set up TGBot4SillyTavern in 5 minutes, even if you've never used Docker or Go before.

## Prerequisites Checklist

Before starting, make sure you have:

- [ ] A computer running Windows, macOS, or Linux
- [ ] [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed (get it from docker.com)
- [ ] A running [SillyTavern](https://github.com/SillyTavern/SillyTavern) instance
- [ ] A Telegram account

## Step 1: Create Your Telegram Bot (5 minutes)

1. Open Telegram on your phone or computer
2. Search for `@BotFather` (the official bot from Telegram)
3. Start a chat and send: `/newbot`
4. Follow the prompts:
   - Choose a name for your bot (e.g., "My SillyTavern Bot")
   - Choose a username (must end in 'bot', e.g., "MySillyTavernBot")
5. **Copy the token** - it looks like: `1234567890:ABCdefGHIjklMNOpqrsTUVwxyz`
6. Save this token somewhere safe - you'll need it in the next step!

## Step 2: Download and Configure (2 minutes)

### Windows

1. Open PowerShell or Command Prompt
2. Run these commands:

```powershell
# Download the project
git clone https://github.com/CambriaDev/TGBot4SillyTavern.git
cd TGBot4SillyTavern

# Run setup
bash setup.sh
```

When prompted, enter:
- Your bot token (from Step 1)
- Your SillyTavern URL (usually `http://localhost:8000`)
- Yes for headless mode
- No for debug mode (unless you want detailed logs)

### macOS / Linux

1. Open Terminal
2. Run these commands:

```bash
# Download the project
git clone https://github.com/CambriaDev/TGBot4SillyTavern.git
cd TGBot4SillyTavern

# Run setup
./setup.sh
```

When prompted, enter:
- Your bot token (from Step 1)
- Your SillyTavern URL (usually `http://localhost:8000`)
- Yes for headless mode
- No for debug mode (unless you want detailed logs)

## Step 3: Start the Bot (1 minute)

Run this command:

```bash
docker-compose up -d
```

Wait for it to download and start (first time takes 1-2 minutes).

## Step 4: Use Your Bot (2 minutes)

1. Open Telegram
2. Search for your bot (the username you chose in Step 1)
3. Send `/start`
4. You'll see a menu with buttons - click them!
5. Select a character, then start chatting!

## Troubleshooting

### "Docker command not found"
- Install Docker Desktop from https://www.docker.com/products/docker-desktop/

### "Cannot connect to SillyTavern"
- Make sure SillyTavern is running
- If using Docker on Windows/Mac, use `http://host.docker.internal:8000` instead of `localhost`
- Edit `.env` file and change `SILLYTAVERN_URL` if needed

### "Bot doesn't respond"
- Check if it's running: `docker-compose logs -f`
- Verify your bot token is correct in `.env`
- Restart: `docker-compose restart`

### "Character list is empty"
- Make sure you have characters in SillyTavern
- Check SillyTavern is accessible at the URL you configured

## Viewing Logs

To see what the bot is doing:

```bash
docker-compose logs -f
```

Press `Ctrl+C` to stop viewing logs (the bot keeps running).

## Stopping the Bot

```bash
docker-compose down
```

## Updating the Bot

```bash
git pull
docker-compose down
docker-compose build
docker-compose up -d
```

## Common Commands Reference

| Command | What it does |
|---------|--------------|
| `docker-compose up -d` | Start the bot in background |
| `docker-compose down` | Stop the bot |
| `docker-compose logs -f` | View bot logs |
| `docker-compose restart` | Restart the bot |
| `docker-compose build` | Rebuild after code changes |

## Getting Help

1. Check [USAGE.md](USAGE.md) for detailed usage instructions
2. Check [README.md](README.md) for technical details
3. Look at [GitHub Issues](https://github.com/CambriaDev/TGBot4SillyTavern/issues)
4. Open a new issue if you're stuck

## Next Steps

Once everything works:

- Explore the `/menu` command for all features
- Try different characters
- Load chat history
- Switch completion presets
- Use the regenerate and edit buttons

Enjoy chatting with your characters on Telegram! 🎉
