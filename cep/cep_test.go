package cep_test

import (
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/cep"
)

func TestIsCEP(t *testing.T) {
	tests := []struct {
		name    string
		cep     string
		wantErr error
	}{
		{"valid CEP with hyphen", "12345-678", nil},
		{"valid CEP without formatting", "12345678", nil},
		{"valid CEP with surrounding spaces", " 12345-678 ", nil},
		{"invalid CEP empty value", "", cep.ErrCEPMustHave8Digits},
		{"invalid CEP with less than 8 digits", "12345-67", cep.ErrCEPMustHave8Digits},
		{"invalid CEP with more than 8 digits", "12345-6789", cep.ErrCEPMustHave8Digits},
		{"invalid CEP with letters", "12345-67A", cep.ErrCEPNotValid},
		{"invalid CEP with unsupported formatting character", "12345.678", cep.ErrCEPNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cep.IsCEP(tt.cep)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsCEP() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsCEPBytes(t *testing.T) {
	tests := []struct {
		name    string
		cep     []byte
		wantErr error
	}{
		{"valid CEP bytes with hyphen", []byte("12345-678"), nil},
		{"valid CEP bytes without formatting", []byte("12345678"), nil},
		{"invalid CEP bytes empty value", []byte(""), cep.ErrCEPMustHave8Digits},
		{"invalid CEP bytes nil value", nil, cep.ErrCEPMustHave8Digits},
		{"invalid CEP bytes with letters", []byte("12345-67A"), cep.ErrCEPNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cep.IsCEPBytes(tt.cep)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsCEPBytes() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
