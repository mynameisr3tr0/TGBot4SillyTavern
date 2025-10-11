package formatter

import (
	"fmt"
	"regexp"
	"strings"
)

// HTMLToTelegram converts HTML text to Telegram-compatible format
func HTMLToTelegram(html string) string {
	text := html

	// Convert <br> and <br/> to newlines
	text = regexp.MustCompile(`<br\s*/?>`).ReplaceAllString(text, "\n")

	// Convert <p> tags to newlines (but keep double newlines between paragraphs)
	text = regexp.MustCompile(`<p[^>]*>`).ReplaceAllString(text, "")
	text = regexp.MustCompile(`</p>`).ReplaceAllString(text, "\n\n")

	// Convert <strong> to Telegram bold <b>
	text = regexp.MustCompile(`<strong[^>]*>`).ReplaceAllString(text, "<b>")
	text = regexp.MustCompile(`</strong>`).ReplaceAllString(text, "</b>")

	// Normalize <b> tags
	text = regexp.MustCompile(`<b[^>]*>`).ReplaceAllString(text, "<b>")

	// Convert <em> to Telegram italic <i>
	text = regexp.MustCompile(`<em[^>]*>`).ReplaceAllString(text, "<i>")
	text = regexp.MustCompile(`</em>`).ReplaceAllString(text, "</i>")

	// Normalize <i> tags
	text = regexp.MustCompile(`<i[^>]*>`).ReplaceAllString(text, "<i>")

	// Normalize <a> tags to Telegram format
	linkRegex := regexp.MustCompile(`<a[^>]*href=["']([^"']*)["'][^>]*>([^<]*)</a>`)
	text = linkRegex.ReplaceAllString(text, `<a href="$1">$2</a>`)

	// Remove all other HTML tags except b, i, and a
	// First, protect our desired tags by replacing them with placeholders
	text = strings.ReplaceAll(text, "<b>", "{{BOLD_OPEN}}")
	text = strings.ReplaceAll(text, "</b>", "{{BOLD_CLOSE}}")
	text = strings.ReplaceAll(text, "<i>", "{{ITALIC_OPEN}}")
	text = strings.ReplaceAll(text, "</i>", "{{ITALIC_CLOSE}}")

	// Protect <a href> tags
	aTagRegex := regexp.MustCompile(`<a href="([^"]*)">(.*?)</a>`)
	matches := aTagRegex.FindAllStringSubmatch(text, -1)
	for i, match := range matches {
		placeholder := fmt.Sprintf("{{LINK_%d}}", i)
		text = strings.Replace(text, match[0], placeholder, 1)
	}

	// Now remove all HTML tags
	text = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(text, "")

	// Restore our desired tags
	text = strings.ReplaceAll(text, "{{BOLD_OPEN}}", "<b>")
	text = strings.ReplaceAll(text, "{{BOLD_CLOSE}}", "</b>")
	text = strings.ReplaceAll(text, "{{ITALIC_OPEN}}", "<i>")
	text = strings.ReplaceAll(text, "{{ITALIC_CLOSE}}", "</i>")

	// Restore links
	for i, match := range matches {
		placeholder := fmt.Sprintf("{{LINK_%d}}", i)
		text = strings.Replace(text, placeholder, fmt.Sprintf(`<a href="%s">%s</a>`, match[1], match[2]), 1)
	}

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
