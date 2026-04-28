package cpf_test

import (
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/cpf"
)

func TestIsCPF(t *testing.T) {
	tests := []struct {
		name    string
		cpf     string
		wantErr error
	}{
		{"valid CPF with dots and dash", "123.456.789-09", nil},
		{"valid CPF without formatting", "12345678909", nil},
		{"valid CPF with first check digit as zero", "39053344705", nil},
		{"valid CPF ignores non digit decorations", "CPF: 123.456.789-09", nil},
		{"invalid CPF with less than 11 digits", "123.456.789-0", cpf.ErrCPFMustHave11Digits},
		{"invalid CPF with more than 11 digits", "123.456.789-090", cpf.ErrCPFMustHave11Digits},
		{"invalid CPF with letters but not enough digits", "1234567890A", cpf.ErrCPFMustHave11Digits},
		{"invalid CPF with invalid first check digit", "12345678992", cpf.ErrCPFNotValid},
		{"invalid CPF with only digits but invalid checksum", "12345678902", cpf.ErrCPFNotValid},
		{"invalid CPF with valid first digit and invalid second digit", "12345678908", cpf.ErrCPFNotValid},
		{"invalid CPF formatted with invalid checksum", "123.456.789-02", cpf.ErrCPFNotValid},
		{"invalid CPF with partial formatting and invalid checksum", "123456789-02", cpf.ErrCPFNotValid},
		{"invalid CPF with repeated zeroes", "00000000000", cpf.ErrCPFNotValid},
		{"invalid CPF with repeated digits", "11111111111", cpf.ErrCPFNotValid},
		{"invalid CPF empty value", "", cpf.ErrCPFMustHave11Digits},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cpf.IsCPF(tt.cpf)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsCPF() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
