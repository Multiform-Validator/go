package base64_test

import (
	stdbase64 "encoding/base64"
	"errors"
	"fmt"
	"strings"
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
		{"valid standard plus character", "++++", nil},
		{"valid standard slash character", "////", nil},
		{"valid url dash character", "----", nil},
		{"valid url underscore character", "____", nil},
		{"valid mixed standard and url alphabet", "+/-_", nil},
		{"valid single byte padded", "Zg==", nil},
		{"valid two bytes padded", "Zm8=", nil},
		{"valid three bytes unpadded", "Zm9v", nil},
		{"valid four bytes padded", "Zm9vYg==", nil},
		{"valid five bytes padded", "Zm9vYmE=", nil},
		{"valid six bytes unpadded", "Zm9vYmFy", nil},
		{"valid raw length two", "Zg", nil},
		{"valid raw length three", "Zm8", nil},
		{"valid raw length four", "Zm9v", nil},
		{"valid uppercase alphabet", "QUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVo=", nil},
		{"valid lowercase alphabet", "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo=", nil},
		{"valid numeric alphabet", "MDEyMzQ1Njc4OQ==", nil},
		{"valid with surrounding tabs", "\tSGVsbG8=\t", nil},
		{"valid with surrounding newlines", "\nSGVsbG8=\r\n", nil},
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
		{"invalid one character", "A", base64.ErrBase64NotValid},
		{"invalid length five", "AAAAA", base64.ErrBase64NotValid},
		{"invalid length nine", "AAAAAAAAA", base64.ErrBase64NotValid},
		{"invalid value with internal space", "SGVs bG8=", base64.ErrBase64NotValid},
		{"invalid value with internal tab", "SGVs\tbG8=", base64.ErrBase64NotValid},
		{"invalid value with line break", "SGVs\nbG8=", base64.ErrBase64NotValid},
		{"invalid value with carriage return", "SGVs\rbG8=", base64.ErrBase64NotValid},
		{"invalid value with form feed", "SGVs\fbG8=", base64.ErrBase64NotValid},
		{"invalid value with vertical tab", "SGVs\vbG8=", base64.ErrBase64NotValid},
		{"invalid base64 characters", "not base64!", base64.ErrBase64NotValid},
		{"invalid unicode character", "SGVsbG8✓", base64.ErrBase64NotValid},
		{"invalid emoji character", "SGVsbG8😀", base64.ErrBase64NotValid},
		{"invalid null byte", "SGVs\x00bG8=", base64.ErrBase64NotValid},
		{"invalid dot character", "SGVsbG8.", base64.ErrBase64NotValid},
		{"invalid comma character", "SGVsbG8,", base64.ErrBase64NotValid},
		{"invalid colon character", "SGVsbG8:", base64.ErrBase64NotValid},
		{"invalid semicolon character", "SGVsbG8;", base64.ErrBase64NotValid},
		{"invalid equals in middle", "SG=VsbG8", base64.ErrBase64NotValid},
		{"invalid equals at start", "=SGVsbG8", base64.ErrBase64NotValid},
		{"invalid padding with short group", "AA=", base64.ErrBase64NotValid},
		{"invalid too much padding only", "====", base64.ErrBase64NotValid},
		{"invalid base64 padding", "SGVsbG8===", base64.ErrBase64NotValid},
		{"invalid base64 impossible length", "A", base64.ErrBase64NotValid},
		{"invalid strict one padding non zero unused bits", "AAB=", base64.ErrBase64NotValid},
		{"invalid strict two padding non zero unused bits", "AB==", base64.ErrBase64NotValid},
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

func TestIsBase64GeneratedFromBinaryData(t *testing.T) {
	inputs := [][]byte{
		{},
		{0},
		{1},
		{255},
		{0, 0},
		{0, 1},
		{255, 254},
		{0, 1, 2},
		{250, 251, 252},
		[]byte("f"),
		[]byte("fo"),
		[]byte("foo"),
		[]byte("foob"),
		[]byte("fooba"),
		[]byte("foobar"),
		[]byte("hello"),
		[]byte("Hello, World!"),
		[]byte("Line 1\nLine 2\r\nLine 3"),
		[]byte(strings.Repeat("a", 64)),
		[]byte(strings.Repeat("Multiform Validator ", 16)),
	}

	encodings := []struct {
		name          string
		encoding      *stdbase64.Encoding
		acceptsEmpty  bool
		requiresInput bool
	}{
		{"standard padded", stdbase64.StdEncoding, false, true},
		{"standard raw", stdbase64.RawStdEncoding, false, true},
		{"url padded", stdbase64.URLEncoding, false, true},
		{"url raw", stdbase64.RawURLEncoding, false, true},
	}

	for _, input := range inputs {
		for _, encoding := range encodings {
			t.Run(fmt.Sprintf("%s %d bytes", encoding.name, len(input)), func(t *testing.T) {
				value := encoding.encoding.EncodeToString(input)
				err := base64.IsBase64(value)
				if len(input) == 0 {
					if !errors.Is(err, base64.ErrBase64NotValid) {
						t.Errorf("IsBase64(%q) error = %v, want %v", value, err, base64.ErrBase64NotValid)
					}
					return
				}

				if err != nil {
					t.Errorf("IsBase64(%q) generated from %v error = %v", value, input, err)
				}
			})
		}
	}
}

func TestIsBase64GeneratedFromAllByteValues(t *testing.T) {
	input := make([]byte, 256)
	for i := range input {
		input[i] = byte(i)
	}

	encodings := []*stdbase64.Encoding{
		stdbase64.StdEncoding,
		stdbase64.RawStdEncoding,
		stdbase64.URLEncoding,
		stdbase64.RawURLEncoding,
	}

	for index, encoding := range encodings {
		t.Run(fmt.Sprintf("encoding %d", index), func(t *testing.T) {
			value := encoding.EncodeToString(input)
			if err := base64.IsBase64(value); err != nil {
				t.Errorf("IsBase64() generated from all byte values error = %v", err)
			}
		})
	}
}

func TestIsBase64StrictPaddingBits(t *testing.T) {
	tests := []struct {
		name    string
		values  []string
		wantErr error
	}{
		{
			name:    "one padding valid last sextets",
			values:  []string{"AAA=", "AAE=", "AAI=", "AAM="},
			wantErr: nil,
		},
		{
			name:    "one padding invalid last sextets",
			values:  []string{"AAB=", "AAF=", "AA/=", "AA_="},
			wantErr: base64.ErrBase64NotValid,
		},
		{
			name:    "two padding valid last sextets",
			values:  []string{"AA==", "AQ==", "Ag==", "Aw==", "BA==", "Dw==", "8A==", "_w=="},
			wantErr: nil,
		},
		{
			name:    "two padding invalid last sextets",
			values:  []string{"AB==", "AC==", "A/==", "A_==", "A-=="},
			wantErr: base64.ErrBase64NotValid,
		},
	}

	for _, tt := range tests {
		for _, value := range tt.values {
			t.Run(tt.name+" "+value, func(t *testing.T) {
				err := base64.IsBase64(value)
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("IsBase64(%q) error = %v, want %v", value, err, tt.wantErr)
				}
			})
		}
	}
}

func TestIsBase64RejectsEveryNonAlphabetASCIICharacter(t *testing.T) {
	for value := byte(0); value < 128; value++ {
		if isAcceptedBase64Byte(value) || value == '=' {
			continue
		}

		t.Run(fmt.Sprintf("byte 0x%02x", value), func(t *testing.T) {
			input := "AA" + string([]byte{value}) + "A"
			err := base64.IsBase64(input)
			if !errors.Is(err, base64.ErrBase64NotValid) {
				t.Errorf("IsBase64(%q) error = %v, want %v", input, err, base64.ErrBase64NotValid)
			}
		})
	}
}

func isAcceptedBase64Byte(value byte) bool {
	return (value >= 'A' && value <= 'Z') ||
		(value >= 'a' && value <= 'z') ||
		(value >= '0' && value <= '9') ||
		value == '+' ||
		value == '/' ||
		value == '-' ||
		value == '_'
}

func TestIsBase64RejectsPaddingInEveryInvalidPosition(t *testing.T) {
	tests := []string{
		"=AAA",
		"A=AA",
		"AA=A",
		"AAAA=AAA",
		"AAA=A==",
		"AA==AA",
		"A===AA",
		"AA===",
		"AAA===",
		"AAAA====",
	}

	for _, value := range tests {
		t.Run(value, func(t *testing.T) {
			err := base64.IsBase64(value)
			if !errors.Is(err, base64.ErrBase64NotValid) {
				t.Errorf("IsBase64(%q) error = %v, want %v", value, err, base64.ErrBase64NotValid)
			}
		})
	}
}

func TestIsBase64LongInputs(t *testing.T) {
	validValue := stdbase64.StdEncoding.EncodeToString([]byte(strings.Repeat("0123456789abcdef", 256)))
	if err := base64.IsBase64(validValue); err != nil {
		t.Errorf("IsBase64(long valid input) error = %v", err)
	}

	invalidValue := validValue[:len(validValue)/2] + "!" + validValue[len(validValue)/2+1:]
	if err := base64.IsBase64(invalidValue); !errors.Is(err, base64.ErrBase64NotValid) {
		t.Errorf("IsBase64(long invalid input) error = %v, want %v", err, base64.ErrBase64NotValid)
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
