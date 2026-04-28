package postalcode_test

import (
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/postalcode"
)

func TestIsPostalCodeGeneric(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr error
	}{
		{"valid united states", "90210", nil},
		{"valid canada", "M4B 1B3", nil},
		{"valid united kingdom", "SW1A 1AA", nil},
		{"valid japan", "100-0001", nil},
		{"valid brazil", "10045-123", nil},
		{"invalid empty", "", postalcode.ErrPostalCodeNotValid},
		{"invalid blank", "   ", postalcode.ErrPostalCodeNotValid},
		{"invalid unsupported format", "ABCDE", postalcode.ErrPostalCodeNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := postalcode.IsPostalCode(tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsPostalCode() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsPostalCodeByCountry(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		country string
		wantErr error
	}{
		{"valid brazil with hyphen", "10045-123", "BR", nil},
		{"valid brazil without hyphen", "10045123", "Brasil", nil},
		{"valid canada with space", "M4B 1B3", "CA", nil},
		{"valid canada without space", "M4B1B3", "Canada", nil},
		{"valid united kingdom", "SW1A 1AA", "UK", nil},
		{"valid united kingdom without space", "SW1A1AA", "United Kingdom", nil},
		{"valid france", "75013", "FR", nil},
		{"valid netherlands letters", "1012 AB", "NL", nil},
		{"valid netherlands numeric", "1012", "Netherlands", nil},
		{"valid japan with hyphen", "100-0001", "JP", nil},
		{"valid japan without hyphen", "1000001", "Japan", nil},
		{"valid spain", "28001", "ES", nil},
		{"valid south africa", "8000", "ZA", nil},
		{"valid germany", "13355", "DE", nil},
		{"valid switzerland", "1002", "CH", nil},
		{"valid italy", "00100", "IT", nil},
		{"valid united states with extension", "90210-1234", "US", nil},
		{"valid united states with blank country uses generic", "90210", "   ", nil},
		{"invalid brazil short", "10045-12", "BR", postalcode.ErrPostalCodeNotValid},
		{"invalid canada format", "123 456", "CA", postalcode.ErrPostalCodeNotValid},
		{"invalid country", "90210", "AR", postalcode.ErrPostalCodeCountryNotSupported},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := postalcode.IsPostalCode(tt.value, tt.country)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsPostalCode() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
