package base64_test

import (
	stdbase64 "encoding/base64"
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/base64"
)

func TestIsBase64(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr error
	}{
		{"valid padded base64", "SGVsbG8=", nil},
		{"valid raw base64", "SGVsbG8", nil},
		{"valid URL base64", "SGVsbG8_", nil},
		{"valid raw URL base64 with dash", "SGVsbG8t", nil},
		{"valid base64 with double padding", "TWE=", nil},
		{"valid base64 with surrounding spaces", " SGVsbG8= ", nil},
		{"valid base64 generated from text with space", "SGVsbG8gV29ybGQ=", nil},
		{"valid base64 generated from text with multiple spaces", "SGVsbG8gICAgV29ybGQ=", nil},
		{"valid base64 generated from text with line feed", "TGluZSAxCkxpbmUgMg==", nil},
		{"valid base64 generated from text with CRLF", "TGluZSAxDQpMaW5lIDI=", nil},
		{"valid base64 generated from text with tab", "Q29sdW1uMQlDb2x1bW4y", nil},
		{"valid base64 generated from text with space and line feed", "SGVsbG8gCldvcmxk", nil},
		{"valid base64 generated from text with blank line", "TGluZSAxCgpMaW5lIDM=", nil},
		{"invalid empty value", "", base64.ErrBase64NotValid},
		{"invalid blank value", " \t\n ", base64.ErrBase64NotValid},
		{"invalid value with internal space", "SGVs bG8=", base64.ErrBase64NotValid},
		{"invalid value with internal tab", "SGVs\tbG8=", base64.ErrBase64NotValid},
		{"invalid value with line break", "SGVs\nbG8=", base64.ErrBase64NotValid},
		{"invalid base64 characters", "not base64!", base64.ErrBase64NotValid},
		{"invalid base64 padding", "SGVsbG8===", base64.ErrBase64NotValid},
		{"invalid base64 impossible length", "A", base64.ErrBase64NotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := base64.IsBase64(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsBase64() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsBase64GeneratedFromWhitespaceContent(t *testing.T) {
	contents := []struct {
		name  string
		value string
	}{
		{"single space", "Hello World"},
		{"leading space", " Hello"},
		{"trailing space", "Hello "},
		{"multiple spaces", "Hello    World"},
		{"line feed", "Line 1\nLine 2"},
		{"carriage return line feed", "Line 1\r\nLine 2"},
		{"tab", "Column1\tColumn2"},
		{"space and line feed", "Hello \nWorld"},
		{"space tab and line feed", "Hello \t\n World"},
		{"blank line", "Line 1\n\nLine 3"},
	}

	encodings := []struct {
		name     string
		encoding *stdbase64.Encoding
	}{
		{"standard padded", stdbase64.StdEncoding},
		{"standard raw", stdbase64.RawStdEncoding},
		{"url padded", stdbase64.URLEncoding},
		{"url raw", stdbase64.RawURLEncoding},
	}

	for _, content := range contents {
		for _, encoding := range encodings {
			t.Run(content.name+" "+encoding.name, func(t *testing.T) {
				value := encoding.encoding.EncodeToString([]byte(content.value))
				err := base64.IsBase64(value)
				if err != nil {
					t.Errorf("IsBase64(%q) generated from %q error = %v", value, content.value, err)
				}
			})
		}
	}
}
