package text_test

import (
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/text"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr error
	}{
		{"valid empty value", "", nil},
		{"invalid value with spaces", "   ", text.ErrValueNotEmpty},
		{"invalid value with null byte", "\x00", text.ErrValueNotEmpty},
		{"invalid value with text", "value", text.ErrValueNotEmpty},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := text.IsEmpty(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsEmpty() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsEmptyBytes(t *testing.T) {
	tests := []struct {
		name    string
		value   []byte
		wantErr error
	}{
		{"valid empty bytes", []byte(""), nil},
		{"valid nil bytes", nil, nil},
		{"invalid bytes with spaces", []byte("   "), text.ErrValueNotEmpty},
		{"invalid bytes with null byte", []byte{0x00}, text.ErrValueNotEmpty},
		{"invalid bytes with text", []byte("value"), text.ErrValueNotEmpty},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := text.IsEmptyBytes(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsEmptyBytes() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr error
	}{
		{"valid empty value", "", nil},
		{"valid value with spaces", "   ", nil},
		{"valid value with tabs and new lines", "\t\n ", nil},
		{"valid value with unicode whitespace", "\u2003\u2009", nil},
		{"invalid value with text", "value", text.ErrValueNotBlank},
		{"invalid value with text and spaces", " value ", text.ErrValueNotBlank},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := text.IsBlank(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsBlank() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsBlankBytes(t *testing.T) {
	tests := []struct {
		name    string
		value   []byte
		wantErr error
	}{
		{"valid empty bytes", []byte(""), nil},
		{"valid nil bytes", nil, nil},
		{"valid bytes with spaces", []byte("   "), nil},
		{"valid bytes with unicode whitespace", []byte("\u2003"), nil},
		{"invalid bytes with null byte", []byte{0x00}, text.ErrValueNotBlank},
		{"invalid bytes with text", []byte("value"), text.ErrValueNotBlank},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := text.IsBlankBytes(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsBlankBytes() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
