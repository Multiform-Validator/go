package base64_test

import (
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
		{"valid base64 with surrounding spaces", " SGVsbG8= ", nil},
		{"invalid empty value", "", base64.ErrBase64NotValid},
		{"invalid value with internal space", "SGVs bG8=", base64.ErrBase64NotValid},
		{"invalid value with line break", "SGVs\nbG8=", base64.ErrBase64NotValid},
		{"invalid base64 characters", "not base64!", base64.ErrBase64NotValid},
		{"invalid base64 padding", "SGVsbG8===", base64.ErrBase64NotValid},
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
