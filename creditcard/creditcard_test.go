package creditcard_test

import (
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/creditcard"
)

func TestIsCreditCard(t *testing.T) {
	tests := []struct {
		name       string
		creditCard string
		wantErr    error
	}{
		{"valid Visa credit card", "4111111111111111", nil},
		{"valid Mastercard credit card", "5555555555554444", nil},
		{"valid Amex credit card", "378282246310005", nil},
		{"valid credit card with spaces", "4111 1111 1111 1111", nil},
		{"valid credit card with hyphens", "4111-1111-1111-1111", nil},
		{"invalid credit card empty value", "", creditcard.ErrCreditCardMustHaveBetween12And19Digits},
		{"invalid credit card with less than 12 digits", "41111111111", creditcard.ErrCreditCardMustHaveBetween12And19Digits},
		{"invalid credit card with more than 19 digits", "41111111111111111111", creditcard.ErrCreditCardMustHaveBetween12And19Digits},
		{"invalid credit card checksum", "4111111111111112", creditcard.ErrCreditCardNotValid},
		{"invalid credit card repeated digits", "0000000000000000", creditcard.ErrCreditCardNotValid},
		{"invalid credit card with unsupported character", "4111.1111.1111.1111", creditcard.ErrCreditCardNotValid},
		{"invalid credit card with letters", "411111111111111A", creditcard.ErrCreditCardNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := creditcard.IsCreditCard(tt.creditCard)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsCreditCard() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsCreditCardBytes(t *testing.T) {
	tests := []struct {
		name       string
		creditCard []byte
		wantErr    error
	}{
		{"valid credit card bytes", []byte("4111111111111111"), nil},
		{"valid credit card bytes with spaces", []byte("4111 1111 1111 1111"), nil},
		{"invalid credit card bytes empty value", []byte(""), creditcard.ErrCreditCardMustHaveBetween12And19Digits},
		{"invalid credit card bytes nil value", nil, creditcard.ErrCreditCardMustHaveBetween12And19Digits},
		{"invalid credit card bytes checksum", []byte("4111111111111112"), creditcard.ErrCreditCardNotValid},
		{"invalid credit card bytes with letters", []byte("411111111111111A"), creditcard.ErrCreditCardNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := creditcard.IsCreditCardBytes(tt.creditCard)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsCreditCardBytes() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
