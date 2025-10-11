#!/bin/bash

# TGBot4SillyTavern Setup Script
# This script helps you configure the bot with interactive prompts

set -e

echo "=========================================="
echo "  TGBot4SillyTavern Setup"
echo "=========================================="
echo ""

# Check if .env already exists
if [ -f .env ]; then
    read -p ".env file already exists. Overwrite? (y/N): " overwrite
    if [[ ! $overwrite =~ ^[Yy]$ ]]; then
        echo "Setup cancelled."
        exit 0
    fi
fi

# Telegram Bot Token
echo "Step 1: Telegram Bot Token"
echo "Get your bot token from @BotFather on Telegram"
read -p "Enter your Telegram Bot Token: " BOT_TOKEN

# SillyTavern URL
echo ""
echo "Step 2: SillyTavern URL"
echo "Default: http://localhost:8000"
read -p "Enter SillyTavern URL (press Enter for default): " ST_URL
ST_URL=${ST_URL:-http://localhost:8000}

# Headless Mode
echo ""
echo "Step 3: Browser Mode"
read -p "Run browser in headless mode? (Y/n): " HEADLESS
if [[ $HEADLESS =~ ^[Nn]$ ]]; then
    HEADLESS_MODE="false"
else
    HEADLESS_MODE="true"
fi

# Debug Mode
echo ""
echo "Step 4: Debug Mode"
read -p "Enable debug logging? (y/N): " DEBUG
if [[ $DEBUG =~ ^[Yy]$ ]]; then
    DEBUG_MODE="true"
else
    DEBUG_MODE="false"
fi

# Write .env file
cat > .env << EOF
# Telegram Bot Configuration
TELEGRAM_BOT_TOKEN=$BOT_TOKEN

# SillyTavern Configuration
SILLYTAVERN_URL=$ST_URL

# Browser Configuration
HEADLESS_MODE=$HEADLESS_MODE
CHROMIUM_PATH=

# Debug Mode
DEBUG=$DEBUG_MODE
EOF

echo ""
echo "=========================================="
echo "  Configuration saved to .env"
echo "=========================================="
echo ""
echo "Your configuration:"
echo "  Bot Token: ${BOT_TOKEN:0:10}..."
echo "  SillyTavern URL: $ST_URL"
echo "  Headless Mode: $HEADLESS_MODE"
echo "  Debug Mode: $DEBUG_MODE"
echo ""
echo "To start the bot:"
echo "  - With Docker: docker-compose up -d"
echo "  - Locally: go run main.go"
echo ""
echo "Check logs:"
echo "  - With Docker: docker-compose logs -f"
echo ""
