package formatter

import (
	"regexp"
	"strings"
)

// HTMLToTelegram converts HTML text to Telegram-compatible format
func HTMLToTelegram(html string) string {
	text := html

	// Convert <br> and <br/> to newlines
	text = regexp.MustCompile(`<br\s*/?>`).ReplaceAllString(text, "\n")
	
	// Convert <p> tags to newlines
	text = regexp.MustCompile(`<p[^>]*>`).ReplaceAllString(text, "")
	text = regexp.MustCompile(`</p>`).ReplaceAllString(text, "\n")

	// Convert <b> and <strong> to Telegram bold
	text = regexp.MustCompile(`<b[^>]*>`).ReplaceAllString(text, "<b>")
	text = regexp.MustCompile(`</b>`).ReplaceAllString(text, "</b>")
	text = regexp.MustCompile(`<strong[^>]*>`).ReplaceAllString(text, "<b>")
	text = regexp.MustCompile(`</strong>`).ReplaceAllString(text, "</b>")

	// Convert <i> and <em> to Telegram italic
	text = regexp.MustCompile(`<i[^>]*>`).ReplaceAllString(text, "<i>")
	text = regexp.MustCompile(`</i>`).ReplaceAllString(text, "</i>")
	text = regexp.MustCompile(`<em[^>]*>`).ReplaceAllString(text, "<i>")
	text = regexp.MustCompile(`</em>`).ReplaceAllString(text, "</i>")

	// Convert <a> tags to Telegram links
	linkRegex := regexp.MustCompile(`<a[^>]*href=["']([^"']*)["'][^>]*>([^<]*)</a>`)
	text = linkRegex.ReplaceAllString(text, `<a href="$1">$2</a>`)

	// Remove other HTML tags
	text = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(text, "")

	// Decode HTML entities
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&quot;", "\"")
	text = strings.ReplaceAll(text, "&#39;", "'")
	text = strings.ReplaceAll(text, "&nbsp;", " ")

	// Clean up multiple newlines
	text = regexp.MustCompile(`\n{3,}`).ReplaceAllString(text, "\n\n")
	
	return strings.TrimSpace(text)
}
