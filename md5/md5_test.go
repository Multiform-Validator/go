package md5_test

import (
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/md5"
)

func TestIsMD5(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr error
	}{
		{"valid lowercase MD5", "d41d8cd98f00b204e9800998ecf8427e", nil},
		{"valid uppercase MD5", "D41D8CD98F00B204E9800998ECF8427E", nil},
		{"invalid MD5 empty value", "", md5.ErrMD5NotValid},
		{"invalid MD5 with less than 32 characters", "d41d8cd98f00b204e9800998ecf8427", md5.ErrMD5NotValid},
		{"invalid MD5 with more than 32 characters", "d41d8cd98f00b204e9800998ecf8427ee", md5.ErrMD5NotValid},
		{"invalid MD5 with unsupported character", "d41d8cd98f00b204e9800998ecf8427g", md5.ErrMD5NotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := md5.IsMD5(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsMD5() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
