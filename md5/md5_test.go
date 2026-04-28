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
		{"valid mixed case MD5", "D41d8CD98f00B204E9800998ecF8427E", nil},
		{"valid MD5 for abc", "900150983cd24fb0d6963f7d28e17f72", nil},
		{"valid MD5 for message digest", "f96b697d7cb7938d525a2f31aaf161d0", nil},
		{"valid MD5 with only zeroes", "00000000000000000000000000000000", nil},
		{"valid MD5 with only nines", "99999999999999999999999999999999", nil},
		{"valid MD5 with lowercase a through f", "abcdefabcdefabcdefabcdefabcdefab", nil},
		{"valid MD5 with uppercase A through F", "ABCDEFABCDEFABCDEFABCDEFABCDEFAB", nil},
		{"valid MD5 with all hex digits repeated", "0123456789abcdef0123456789abcdef", nil},
		{"invalid MD5 empty value", "", md5.ErrMD5NotValid},
		{"invalid MD5 blank value", "                                ", md5.ErrMD5NotValid},
		{"invalid MD5 with less than 32 characters", "d41d8cd98f00b204e9800998ecf8427", md5.ErrMD5NotValid},
		{"invalid MD5 with more than 32 characters", "d41d8cd98f00b204e9800998ecf8427ee", md5.ErrMD5NotValid},
		{"invalid MD5 with unsupported character", "d41d8cd98f00b204e9800998ecf8427g", md5.ErrMD5NotValid},
		{"invalid MD5 with surrounding space", " d41d8cd98f00b204e9800998ecf8427e", md5.ErrMD5NotValid},
		{"invalid MD5 with trailing space", "d41d8cd98f00b204e9800998ecf8427e ", md5.ErrMD5NotValid},
		{"invalid MD5 with internal space", "d41d8cd98f00b204 e9800998ecf8427e", md5.ErrMD5NotValid},
		{"invalid MD5 with newline", "d41d8cd98f00b204e9800998ecf8427\n", md5.ErrMD5NotValid},
		{"invalid MD5 with tab", "d41d8cd98f00b204e9800998ecf8427\t", md5.ErrMD5NotValid},
		{"invalid MD5 with internal hyphen", "d41d8cd98f00b204e9800998ecf-427", md5.ErrMD5NotValid},
		{"invalid MD5 with colon separators", "d41d8cd9:8f00b204:e9800998:ecf8427e", md5.ErrMD5NotValid},
		{"invalid MD5 with punctuation", "d41d8cd98f00b204e9800998ecf8427!", md5.ErrMD5NotValid},
		{"invalid MD5 with unicode accent", "d41d8cd98f00b204e9800998ecf842é", md5.ErrMD5NotValid},
		{"invalid MD5 with fullwidth hex-like character", "d41d8cd98f00b204e9800998ecf842Ａ", md5.ErrMD5NotValid},
		{"invalid MD5 with negative sign prefix", "-41d8cd98f00b204e9800998ecf8427e", md5.ErrMD5NotValid},
		{"invalid MD5 with 0x prefix", "0xd41d8cd98f00b204e9800998ecf8427e", md5.ErrMD5NotValid},
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
