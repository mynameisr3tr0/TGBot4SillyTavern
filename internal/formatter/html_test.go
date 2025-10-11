package formatter

import (
	"testing"
)

func TestHTMLToTelegram(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Bold tag",
			input:    "<b>Hello World</b>",
			expected: "<b>Hello World</b>",
		},
		{
			name:     "Strong tag to bold",
			input:    "<strong>Hello World</strong>",
			expected: "<b>Hello World</b>",
		},
		{
			name:     "Italic tag",
			input:    "<i>Hello World</i>",
			expected: "<i>Hello World</i>",
		},
		{
			name:     "Em tag to italic",
			input:    "<em>Hello World</em>",
			expected: "<i>Hello World</i>",
		},
		{
			name:     "Link tag",
			input:    `<a href="https://example.com">Click here</a>`,
			expected: `<a href="https://example.com">Click here</a>`,
		},
		{
			name:     "Line break",
			input:    "Hello<br>World",
			expected: "Hello\nWorld",
		},
		{
			name:     "Paragraph tags",
			input:    "<p>Hello</p><p>World</p>",
			expected: "Hello\n\nWorld",
		},
		{
			name:     "HTML entities",
			input:    "Hello &amp; World &lt;test&gt;",
			expected: "Hello & World <test>",
		},
		{
			name:     "Mixed formatting",
			input:    "<p><b>Hello</b> <i>World</i></p>",
			expected: "<b>Hello</b> <i>World</i>",
		},
		{
			name:     "Remove unknown tags",
			input:    "<div>Hello</div> <span>World</span>",
			expected: "Hello World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HTMLToTelegram(tt.input)
			if result != tt.expected {
				t.Errorf("HTMLToTelegram() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestHTMLToTelegramMultipleNewlines(t *testing.T) {
	input := "Hello\n\n\n\n\nWorld"
	expected := "Hello\n\nWorld"
	result := HTMLToTelegram(input)
	if result != expected {
		t.Errorf("HTMLToTelegram() = %q, want %q", result, expected)
	}
}

func TestHTMLToTelegramTrimWhitespace(t *testing.T) {
	input := "  <p>Hello World</p>  "
	expected := "Hello World"
	result := HTMLToTelegram(input)
	if result != expected {
		t.Errorf("HTMLToTelegram() = %q, want %q", result, expected)
	}
}
