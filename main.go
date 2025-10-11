package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/CambriaDev/TGBot4SillyTavern/config"
	"github.com/CambriaDev/TGBot4SillyTavern/internal/bot"
	"github.com/CambriaDev/TGBot4SillyTavern/internal/browser"
)

func main() {
	// Load configuration
	cfg := config.Load()

	if cfg.TelegramToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	// Initialize browser
	log.Println("Initializing headless browser...")
	br, err := browser.New(cfg.SillyTavernURL, cfg.HeadlessMode)
	if err != nil {
		log.Fatalf("Failed to initialize browser: %v", err)
	}
	defer br.Close()

	// Initialize bot
	log.Println("Initializing Telegram bot...")
	b, err := bot.New(cfg.TelegramToken, br)
	if err != nil {
		log.Fatalf("Failed to initialize bot: %v", err)
	}

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Shutting down...")
		br.Close()
		os.Exit(0)
	}()

	// Start bot
	log.Println("Bot started successfully!")
	b.Start()
}
