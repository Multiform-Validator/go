package telephone_test

import (
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/telephone"
)

func TestIsTelephoneGeneric(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr error
	}{
		{"valid international number", "+55 11 91234-5678", nil},
		{"valid local formatted number", "(11) 91234-5678", nil},
		{"valid with dots", "212.555.0199", nil},
		{"valid with tabs and newlines", "212\t555\n0199", nil},
		{"invalid empty value", "", telephone.ErrTelephoneNotValid},
		{"invalid blank value", "   ", telephone.ErrTelephoneNotValid},
		{"invalid short value", "123456", telephone.ErrTelephoneNotValid},
		{"invalid long value", "+1234567890123456", telephone.ErrTelephoneNotValid},
		{"invalid letters", "+55 11 phone", telephone.ErrTelephoneNotValid},
		{"invalid plus position", "55+11912345678", telephone.ErrTelephoneNotValid},
		{"invalid repeated plus", "++55 11 91234-5678", telephone.ErrTelephoneNotValid},
		{"invalid slash separator", "+55/11/91234-5678", telephone.ErrTelephoneNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := telephone.IsTelephone(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsTelephone() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsTelephoneByCountry(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		country string
		wantErr error
	}{
		{"valid brazil mobile with iso country", "+55 11 91234-5678", "BR", nil},
		{"valid brazil mobile with numeric country alias", "+55 11 91234-5678", "+55", nil},
		{"valid brazil landline with country name", "(11) 3456-7890", "Brasil", nil},
		{"valid brazil strips country code without plus when longer than national max", "5511912345678", "BR", nil},
		{"valid united states", "+1 (212) 555-0199", "US", nil},
		{"valid united states with padded country name", "+1 (212) 555-0199", " United   States ", nil},
		{"valid canada", "+1 416 555 0199", "Canada", nil},
		{"valid china", "+86 131 2345 6789", "CN", nil},
		{"valid japan", "+81 90 1234 5678", "JP", nil},
		{"valid germany", "+49 30 123456", "DE", nil},
		{"valid india", "+91 98765 43210", "IN", nil},
		{"valid united kingdom", "+44 20 7946 0958", "UK", nil},
		{"valid france", "+33 1 42 68 53 00", "FR", nil},
		{"valid italy", "+39 06 6982", "IT", nil},
		{"valid south korea", "+82 10 1234 5678", "KR", nil},
		{"invalid brazil area code", "+55 01 91234-5678", "BR", telephone.ErrTelephoneNotValid},
		{"invalid brazil too short after country code strip", "+55 11 1234", "BR", telephone.ErrTelephoneNotValid},
		{"invalid united states area code", "+1 112 555 0199", "US", telephone.ErrTelephoneNotValid},
		{"invalid france prefix zero", "+33 0 42 68 53 00", "FR", telephone.ErrTelephoneNotValid},
		{"invalid china prefix zero", "+86 031 2345 6789", "CN", telephone.ErrTelephoneNotValid},
		{"invalid india mobile prefix", "+91 51234 56789", "IN", telephone.ErrTelephoneNotValid},
		{"invalid country", "+55 11 91234-5678", "AR", telephone.ErrTelephoneCountryNotSupported},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := telephone.IsTelephone(tt.value, tt.country)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsTelephone() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
