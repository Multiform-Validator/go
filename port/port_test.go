package port_test

import (
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/port"
)

func TestIsPort(t *testing.T) {
	tests := []struct {
		name    string
		port    string
		wantErr error
	}{
		{"valid minimum port", "1", nil},
		{"valid common port", "8080", nil},
		{"valid maximum port", "65535", nil},
		{"valid port with leading zeroes", "00080", nil},
		{"valid port with surrounding spaces", " 443 ", nil},
		{"valid port with surrounding newline", "\n443\t", nil},
		{"invalid port empty value", "", port.ErrPortNotValid},
		{"invalid port zero", "0", port.ErrPortMustBeBetween1And65535},
		{"invalid port all zeroes", "00000", port.ErrPortMustBeBetween1And65535},
		{"invalid port above maximum", "65536", port.ErrPortMustBeBetween1And65535},
		{"invalid port huge numeric value", "999999999999999999999999999999", port.ErrPortMustBeBetween1And65535},
		{"invalid port with letters", "80a", port.ErrPortNotValid},
		{"invalid port with sign", "+80", port.ErrPortNotValid},
		{"invalid port with decimal point", "80.5", port.ErrPortNotValid},
		{"invalid port with internal whitespace", "8 0", port.ErrPortNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := port.IsPort(tt.port)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsPort() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsPortNumber(t *testing.T) {
	tests := []struct {
		name    string
		port    int
		wantErr error
	}{
		{"valid minimum port number", 1, nil},
		{"valid common port number", 8080, nil},
		{"valid maximum port number", 65535, nil},
		{"valid privileged port number", 22, nil},
		{"invalid port number zero", 0, port.ErrPortMustBeBetween1And65535},
		{"invalid port number negative", -1, port.ErrPortMustBeBetween1And65535},
		{"invalid port number above maximum", 65536, port.ErrPortMustBeBetween1And65535},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := port.IsPortNumber(tt.port)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsPortNumber() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
