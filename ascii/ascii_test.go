package ascii_test

import (
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/ascii"
)

func TestIsAscii(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr error
	}{
		{"valid ASCII value", "Hello 123!", nil},
		{"valid empty value", "", nil},
		{"valid ASCII control characters", "\x00\t\n\r\x7f", nil},
		{"invalid value with accent", "olá", ascii.ErrASCIINotValid},
		{"invalid value with emoji", "hello 🙂", ascii.ErrASCIINotValid},
		{"invalid replacement character", "\ufffd", ascii.ErrASCIINotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ascii.IsAscii(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsAscii() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsAsciiBytes(t *testing.T) {
	tests := []struct {
		name    string
		value   []byte
		wantErr error
	}{
		{"valid ASCII bytes", []byte("Hello 123!"), nil},
		{"valid empty bytes", []byte(""), nil},
		{"valid nil bytes", nil, nil},
		{"valid ASCII boundary bytes", []byte{0x00, 0x09, 0x0A, 0x7F}, nil},
		{"invalid bytes with non ASCII value", []byte{0x48, 0x80}, ascii.ErrASCIINotValid},
		{"invalid bytes with UTF-8 accent", []byte("olá"), ascii.ErrASCIINotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ascii.IsAsciiBytes(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsAsciiBytes() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
