package cpf_test

import (
	"testing"

	"github.com/Multiform-Validator/go/cpf"
)

func TestIsCPFValid(t *testing.T) {
	tests := []struct {
		name    string
		cpf     string
		wantErr bool
		wantMsg string
	}{
		{"valid CPF with dots and dash", "123.456.789-09", false, ""},
		{"valid CPF without formatting", "12345678909", false, ""},
		{"valid CPF with first check digit as zero", "39053344705", false, ""},
		{"invalid CPF with less than 11 digits", "123.456.789-0", true, "CPF must have 11 digits"},
		{"invalid CPF with more than 11 digits", "123.456.789-090", true, "CPF must have 11 digits"},
		{"invalid CPF with invalid first check digit", "12345678992", true, "CPF is not valid"},
		{"invalid CPF with only digits but invalid checksum", "12345678902", true, "CPF is not valid"},
		{"invalid CPF with valid first digit and invalid second digit", "12345678908", true, "CPF is not valid"},
		{"invalid CPF formatted with invalid checksum", "123.456.789-02", true, "CPF is not valid"},
		{"invalid CPF with partial formatting and invalid checksum", "123456789-02", true, "CPF is not valid"},
		{"invalid CPF with repeated digits", "11111111111", true, "CPF is not valid"},
		{"invalid CPF empty value", "", true, "CPF must have 11 digits"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cpf.IsCPFValid(tt.cpf)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsCPFValid() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && err != nil && err.Error() != tt.wantMsg {
				t.Errorf("IsCPFValid() error message = %q, want %q", err.Error(), tt.wantMsg)
			}
		})
	}
}

func TestIsCPFValidBytes(t *testing.T) {
	tests := []struct {
		name    string
		cpf     []byte
		wantErr bool
		wantMsg string
	}{
		{"valid CPF bytes with dots and dash", []byte("123.456.789-09"), false, ""},
		{"valid CPF bytes without formatting", []byte("12345678909"), false, ""},
		{"valid CPF bytes with first check digit as zero", []byte("39053344705"), false, ""},
		{"invalid CPF bytes with less than 11 digits", []byte("123.456.789-0"), true, "CPF must have 11 digits"},
		{"invalid CPF bytes with more than 11 digits", []byte("123.456.789-090"), true, "CPF must have 11 digits"},
		{"invalid CPF bytes with invalid first check digit", []byte("12345678992"), true, "CPF is not valid"},
		{"invalid CPF bytes with invalid checksum", []byte("12345678902"), true, "CPF is not valid"},
		{"invalid CPF bytes with valid first digit and invalid second digit", []byte("12345678908"), true, "CPF is not valid"},
		{"invalid CPF bytes with repeated digits", []byte("11111111111"), true, "CPF is not valid"},
		{"invalid CPF bytes empty value", []byte(""), true, "CPF must have 11 digits"},
		{"invalid CPF bytes nil value", nil, true, "CPF must have 11 digits"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cpf.IsCPFValidBytes(tt.cpf)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsCPFValidBytes() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && err != nil && err.Error() != tt.wantMsg {
				t.Errorf("IsCPFValidBytes() error message = %q, want %q", err.Error(), tt.wantMsg)
			}
		})
	}
}
