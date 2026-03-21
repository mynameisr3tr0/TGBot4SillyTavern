package browser

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

// Browser manages the headless browser instance
type Browser struct {
	ctx    context.Context
	cancel context.CancelFunc
	url    string
}

// New creates a new browser instance
func New(sillyTavernURL string, headless bool) (*Browser, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", headless),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
	)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	b := &Browser{
		ctx:    ctx,
		cancel: cancel,
		url:    sillyTavernURL,
	}

	// Navigate to SillyTavern
	if err := chromedp.Run(b.ctx, chromedp.Navigate(sillyTavernURL)); err != nil {
		cancel()
		return nil, fmt.Errorf("failed to navigate to SillyTavern: %w", err)
	}

	// Wait for page to load
	time.Sleep(3 * time.Second)

	return b, nil
}

// Close closes the browser
func (b *Browser) Close() {
	if b.cancel != nil {
		b.cancel()
	}
}

// GetCharacters retrieves the list of available characters
func (b *Browser) GetCharacters() ([]Character, error) {
	var characters []Character

	// This is a placeholder - actual implementation would interact with SillyTavern's DOM
	// to extract character list
	err := chromedp.Run(b.ctx,
		chromedp.WaitVisible(`#rm_ch_create_block`, chromedp.ByID),
		chromedp.Evaluate(`
			Array.from(document.querySelectorAll('.character_select')).map(el => ({
				name: el.querySelector('.ch_name')?.textContent || '',
				avatar: el.querySelector('img')?.src || ''
			}))
		`, &characters),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get characters: %w", err)
	}

	return characters, nil
}

// SelectCharacter selects a character in SillyTavern
func (b *Browser) SelectCharacter(characterName string) error {
	// Click on the character element
	selector := fmt.Sprintf(`//div[contains(@class, 'character_select')]//div[@class='ch_name' and text()='%s']`, characterName)

	err := chromedp.Run(b.ctx,
		chromedp.Click(selector, chromedp.BySearch),
		chromedp.Sleep(1*time.Second),
	)

	if err != nil {
		return fmt.Errorf("failed to select character: %w", err)
	}

	return nil
}

// GetChatHistory retrieves chat history list
func (b *Browser) GetChatHistory() ([]ChatHistory, error) {
	var histories []ChatHistory

	err := chromedp.Run(b.ctx,
		chromedp.Click(`#option_select_chat`, chromedp.ByID),
		chromedp.Sleep(500*time.Millisecond),
		chromedp.Evaluate(`
			Array.from(document.querySelectorAll('.select_chat_block_item')).map(el => ({
				name: el.textContent.trim(),
				timestamp: el.dataset.timestamp || ''
			}))
		`, &histories),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get chat history: %w", err)
	}

	return histories, nil
}

// SelectChatHistory loads a specific chat history
func (b *Browser) SelectChatHistory(historyName string) error {
	selector := fmt.Sprintf(`//div[@class='select_chat_block_item' and contains(text(), '%s')]`, historyName)

	err := chromedp.Run(b.ctx,
		chromedp.Click(`#option_select_chat`, chromedp.ByID),
		chromedp.Sleep(500*time.Millisecond),
		chromedp.Click(selector, chromedp.BySearch),
		chromedp.Sleep(1*time.Second),
	)

	if err != nil {
		return fmt.Errorf("failed to select chat history: %w", err)
	}

	return nil
}

// SendMessage sends a message in the chat and waits for the full response
func (b *Browser) SendMessage(message string) (string, error) {
	var response string

	err := chromedp.Run(b.ctx,
		// 1. Вбиваем текст и жмем отправить
		chromedp.WaitVisible(`#send_textarea`, chromedp.ByID),
		chromedp.SetValue(`#send_textarea`, message, chromedp.ByID),
		chromedp.Click(`#send_but`, chromedp.ByID),
		
		// 2. УМНОЕ ОЖИДАНИЕ: Ждем, пока появится кнопка "Stop" (начало генерации)
		chromedp.WaitVisible(`#mes_stop`, chromedp.ByID),
		
		// 3. УМНОЕ ОЖИДАНИЕ: Ждем, пока кнопка "Stop" ИСЧЕЗНЕТ (конец генерации).
		// Это может занять хоть 5 секунд, хоть минуту — бот будет ждать.
		chromedp.WaitNotVisible(`#mes_stop`, chromedp.ByID),
		
		// Небольшая пауза, чтобы DOM точно успел обновить текст после остановки
		chromedp.Sleep(500*time.Millisecond),
		
		// 4. Забираем текст самого последнего сообщения в чате
		chromedp.Evaluate(`
			(() => {
				const messages = Array.from(document.querySelectorAll('.mes_text'));
				if (messages.length === 0) return "Ошибка: сообщения не найдены";
				return messages[messages.length - 1].innerText;
			})()
		`, &response),
	)

	if err != nil {
		fmt.Printf("DEBUG ERROR in SendMessage: %v\n", err)
		return "", err
	}

	return response, nil
}
// EditMessage edits a specific message
func (b *Browser) EditMessage(messageIndex int, newText string) error {
	script := fmt.Sprintf(`
		const messages = document.querySelectorAll('.mes_block');
		if (messages[%d]) {
			messages[%d].querySelector('.mes_edit')?.click();
		}
	`, messageIndex, messageIndex)

	err := chromedp.Run(b.ctx,
		chromedp.Evaluate(script, nil),
		chromedp.Sleep(500*time.Millisecond),
		chromedp.SetValue(`.edit_textarea`, newText, chromedp.ByQuery),
		chromedp.Click(`.mes_edit_done`, chromedp.ByQuery),
		chromedp.Sleep(1*time.Second),
	)

	if err != nil {
		return fmt.Errorf("failed to edit message: %w", err)
	}

	return nil
}

// RegenerateMessage regenerates the last AI response
func (b *Browser) RegenerateMessage() (string, error) {
	var response string

	err := chromedp.Run(b.ctx,
		chromedp.Click(`#option_regenerate`, chromedp.ByID),
		chromedp.Sleep(2*time.Second),
		chromedp.Evaluate(`
			document.querySelector('.mes_block:last-child .mes_text')?.innerHTML || ''
		`, &response),
	)

	if err != nil {
		return "", fmt.Errorf("failed to regenerate message: %w", err)
	}

	return response, nil
}

// GetPresets retrieves available completion presets
func (b *Browser) GetPresets() ([]Preset, error) {
	var presets []Preset

	err := chromedp.Run(b.ctx,
		chromedp.Evaluate(`
			Array.from(document.querySelectorAll('#settings_preset option')).map(el => ({
				name: el.textContent,
				value: el.value
			}))
		`, &presets),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get presets: %w", err)
	}

	return presets, nil
}

// SelectPreset selects a completion preset
func (b *Browser) SelectPreset(presetName string) error {
	err := chromedp.Run(b.ctx,
		chromedp.SetValue(`#settings_preset`, presetName, chromedp.ByID),
		chromedp.Sleep(500*time.Millisecond),
	)

	if err != nil {
		return fmt.Errorf("failed to select preset: %w", err)
	}

	return nil
}

// Character represents a SillyTavern character
type Character struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// ChatHistory represents a chat history entry
type ChatHistory struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
}

// Preset represents a completion preset
type Preset struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
