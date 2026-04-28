package macaddress_test

import (
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/macaddress"
)

func TestIsMACAddress(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr error
	}{
		{"valid plain", "001A2B3C4D5E", nil},
		{"valid colon separated", "00:1A:2B:3C:4D:5E", nil},
		{"valid hyphen separated", "00-1A-2B-3C-4D-5E", nil},
		{"valid dot separated", "001A.2B3C.4D5E", nil},
		{"valid lowercase", "00:1a:2b:3c:4d:5e", nil},
		{"invalid empty", "", macaddress.ErrMACAddressNotValid},
		{"invalid short", "00:1A:2B:3C:4D", macaddress.ErrMACAddressNotValid},
		{"invalid long", "00:1A:2B:3C:4D:5E:6F", macaddress.ErrMACAddressNotValid},
		{"invalid non hex alpha", "00:1A:2B:3C:4D:ZZ", macaddress.ErrMACAddressNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := macaddress.IsMACAddress(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsMACAddress() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
