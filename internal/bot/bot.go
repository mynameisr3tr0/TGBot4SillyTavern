package bot

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/CambriaDev/TGBot4SillyTavern/internal/browser"
	"github.com/CambriaDev/TGBot4SillyTavern/internal/formatter"
)

// Bot represents the Telegram bot
type Bot struct {
	api     *tgbotapi.BotAPI
	browser *browser.Browser
	updates tgbotapi.UpdatesChannel
}

// New creates a new Telegram bot
func New(token string, br *browser.Browser) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	log.Printf("Authorized on account %s", api.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := api.GetUpdatesChan(u)

	return &Bot{
		api:     api,
		browser: br,
		updates: updates,
	}, nil
}

// Start starts the bot
func (b *Bot) Start() {
	for update := range b.updates {
		if update.Message != nil {
			b.handleMessage(update.Message)
		} else if update.CallbackQuery != nil {
			b.handleCallback(update.CallbackQuery)
		}
	}
}

// handleMessage handles incoming messages
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	if message.IsCommand() {
		b.handleCommand(message)
		return
	}

	// Send user message to SillyTavern and get response
	response, err := b.browser.SendMessage(message.Text)
	if err != nil {
		log.Printf("Error sending message: %v", err)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Failed to send message to SillyTavern")
		b.api.Send(msg)
		return
	}

	// Convert HTML to Telegram format
	formattedResponse := formatter.HTMLToTelegram(response)

	// Create inline keyboard with action buttons
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔄 Regenerate", "regenerate"),
			tgbotapi.NewInlineKeyboardButtonData("✏️ Edit", "edit_last"),
		),
	)

	msg := tgbotapi.NewMessage(message.Chat.ID, formattedResponse)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = keyboard
	b.api.Send(msg)
}

// handleCommand handles bot commands
func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		b.sendMainMenu(message.Chat.ID)
	case "menu":
		b.sendMainMenu(message.Chat.ID)
	case "characters":
		b.showCharacters(message.Chat.ID)
	case "history":
		b.showHistory(message.Chat.ID, 0)
	case "presets":
		b.showPresets(message.Chat.ID)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command. Use /menu to see available options.")
		b.api.Send(msg)
	}
}

// sendMainMenu sends the main menu
func (b *Bot) sendMainMenu(chatID int64) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("👤 Select Character", "show_characters"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💬 Chat History", "show_history:0"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⚙️ Completion Presets", "show_presets"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Welcome to SillyTavern Bot!\n\nChoose an option:")
	msg.ReplyMarkup = keyboard
	b.api.Send(msg)
}

// showCharacters displays character selection
func (b *Bot) showCharacters(chatID int64) {
	characters, err := b.browser.GetCharacters()
	if err != nil {
		log.Printf("Error getting characters: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Failed to load characters")
		b.api.Send(msg)
		return
	}

	var rows [][]tgbotapi.InlineKeyboardButton
	for _, char := range characters {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(char.Name, "char:"+char.Name),
		)
		rows = append(rows, row)
	}
	
	// Add back button
	rows = append(rows, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("« Back to Menu", "main_menu"),
	))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)
	msg := tgbotapi.NewMessage(chatID, "Select a character:")
	msg.ReplyMarkup = keyboard
	b.api.Send(msg)
}

// showHistory displays chat history with pagination
func (b *Bot) showHistory(chatID int64, page int) {
	histories, err := b.browser.GetChatHistory()
	if err != nil {
		log.Printf("Error getting chat history: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Failed to load chat history")
		b.api.Send(msg)
		return
	}

	pageSize := 5
	start := page * pageSize
	end := start + pageSize
	if end > len(histories) {
		end = len(histories)
	}

	var rows [][]tgbotapi.InlineKeyboardButton
	for i := start; i < end; i++ {
		hist := histories[i]
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(hist.Name, "hist:"+hist.Name),
		)
		rows = append(rows, row)
	}

	// Pagination buttons
	var navButtons []tgbotapi.InlineKeyboardButton
	if page > 0 {
		navButtons = append(navButtons, tgbotapi.NewInlineKeyboardButtonData("« Previous", fmt.Sprintf("show_history:%d", page-1)))
	}
	if end < len(histories) {
		navButtons = append(navButtons, tgbotapi.NewInlineKeyboardButtonData("Next »", fmt.Sprintf("show_history:%d", page+1)))
	}
	if len(navButtons) > 0 {
		rows = append(rows, navButtons)
	}

	// Back button
	rows = append(rows, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("« Back to Menu", "main_menu"),
	))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Chat History (Page %d):", page+1))
	msg.ReplyMarkup = keyboard
	b.api.Send(msg)
}

// showPresets displays completion presets
func (b *Bot) showPresets(chatID int64) {
	presets, err := b.browser.GetPresets()
	if err != nil {
		log.Printf("Error getting presets: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Failed to load presets")
		b.api.Send(msg)
		return
	}

	var rows [][]tgbotapi.InlineKeyboardButton
	for _, preset := range presets {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(preset.Name, "preset:"+preset.Name),
		)
		rows = append(rows, row)
	}

	// Add back button
	rows = append(rows, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("« Back to Menu", "main_menu"),
	))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)
	msg := tgbotapi.NewMessage(chatID, "Select a completion preset:")
	msg.ReplyMarkup = keyboard
	b.api.Send(msg)
}

// handleCallback handles callback queries from inline keyboards
func (b *Bot) handleCallback(callback *tgbotapi.CallbackQuery) {
	// Answer callback to remove loading state
	b.api.Request(tgbotapi.NewCallback(callback.ID, ""))

	data := callback.Data
	chatID := callback.Message.Chat.ID

	switch {
	case data == "main_menu":
		b.sendMainMenu(chatID)
	case data == "show_characters":
		b.showCharacters(chatID)
	case strings.HasPrefix(data, "show_history:"):
		pageStr := strings.TrimPrefix(data, "show_history:")
		page, _ := strconv.Atoi(pageStr)
		b.showHistory(chatID, page)
	case data == "show_presets":
		b.showPresets(chatID)
	case strings.HasPrefix(data, "char:"):
		charName := strings.TrimPrefix(data, "char:")
		b.selectCharacter(chatID, charName)
	case strings.HasPrefix(data, "hist:"):
		histName := strings.TrimPrefix(data, "hist:")
		b.selectHistory(chatID, histName)
	case strings.HasPrefix(data, "preset:"):
		presetName := strings.TrimPrefix(data, "preset:")
		b.selectPreset(chatID, presetName)
	case data == "regenerate":
		b.regenerateResponse(chatID)
	case data == "edit_last":
		msg := tgbotapi.NewMessage(chatID, "To edit the message, please send the new text:")
		b.api.Send(msg)
	}
}

// selectCharacter selects a character
func (b *Bot) selectCharacter(chatID int64, charName string) {
	err := b.browser.SelectCharacter(charName)
	if err != nil {
		log.Printf("Error selecting character: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Failed to select character")
		b.api.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("✅ Selected character: %s", charName))
	b.api.Send(msg)
}

// selectHistory loads a chat history
func (b *Bot) selectHistory(chatID int64, histName string) {
	err := b.browser.SelectChatHistory(histName)
	if err != nil {
		log.Printf("Error selecting chat history: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Failed to load chat history")
		b.api.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("✅ Loaded chat: %s", histName))
	b.api.Send(msg)
}

// selectPreset selects a completion preset
func (b *Bot) selectPreset(chatID int64, presetName string) {
	err := b.browser.SelectPreset(presetName)
	if err != nil {
		log.Printf("Error selecting preset: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Failed to select preset")
		b.api.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("✅ Selected preset: %s", presetName))
	b.api.Send(msg)
}

// regenerateResponse regenerates the last AI response
func (b *Bot) regenerateResponse(chatID int64) {
	response, err := b.browser.RegenerateMessage()
	if err != nil {
		log.Printf("Error regenerating message: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Failed to regenerate message")
		b.api.Send(msg)
		return
	}

	formattedResponse := formatter.HTMLToTelegram(response)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔄 Regenerate", "regenerate"),
			tgbotapi.NewInlineKeyboardButtonData("✏️ Edit", "edit_last"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, formattedResponse)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = keyboard
	b.api.Send(msg)
}
