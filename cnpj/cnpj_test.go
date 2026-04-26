package cnpj

import (
	"errors"
	"testing"
)

func TestIsCNPJ(t *testing.T) {
	tests := []struct {
		name    string
		cnpj    string
		wantErr error
	}{
		{"valid alphanumeric CNPJ with formatting", "12.ABC.345/01DE-35", nil},
		{"valid alphanumeric CNPJ without formatting", "12ABC34501DE35", nil},
		{"valid alphanumeric CNPJ with surrounding spaces", " 12.ABC.345/01DE-35 ", nil},
		{"valid numeric CNPJ with formatting", "04.252.011/0001-10", nil},
		{"valid numeric CNPJ without formatting", "11222333000181", nil},
		{"valid numeric CNPJ with alternative formatting", "69.228.768.0159-00", nil},
		{"valid numeric CNPJ corrected from legacy fixture", "72.501.263/0001-95", nil},
		{"valid numeric CNPJ with first check digit as zero", "00000000000604", nil},
		{"invalid CNPJ with less than 14 characters", "12.ABC.345/01DE-3", ErrCNPJMustHave14Characters},
		{"invalid CNPJ with more than 14 characters", "12.ABC.345/01DE-350", ErrCNPJMustHave14Characters},
		{"invalid CNPJ with invalid checksum", "12ABC34501DE34", ErrCNPJNotValid},
		{"invalid numeric CNPJ with invalid checksum", "12.345.678/0001-91", ErrCNPJNotValid},
		{"invalid numeric CNPJ with invalid verifier digits", "12.345.678/0001-00", ErrCNPJNotValid},
		{"invalid numeric CNPJ from legacy fixture", "72.501.263/0001-40", ErrCNPJNotValid},
		{"invalid CNPJ with lowercase letters", "12abc34501DE35", ErrCNPJNotValid},
		{"invalid CNPJ with letter as check digit", "12ABC34501DEA5", ErrCNPJNotValid},
		{"invalid numeric CNPJ with letter as check digit", "72.501.263/0001-4A", ErrCNPJNotValid},
		{"invalid CNPJ with unsupported formatting character", "12ABC34501D@35", ErrCNPJNotValid},
		{"invalid CNPJ with only zeroes", "00000000000000", ErrCNPJNotValid},
		{"invalid numeric CNPJ with all digits repeated", "11.111.111/1111-11", ErrCNPJNotValid},
		{"invalid CNPJ empty value", "", ErrCNPJMustHave14Characters},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsCNPJ(tt.cnpj)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsCNPJ() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsCNPJBytes(t *testing.T) {
	tests := []struct {
		name    string
		cnpj    []byte
		wantErr error
	}{
		{"valid CNPJ bytes with formatting", []byte("12.ABC.345/01DE-35"), nil},
		{"valid CNPJ bytes without formatting", []byte("12ABC34501DE35"), nil},
		{"valid numeric CNPJ bytes with formatting", []byte("04.252.011/0001-10"), nil},
		{"valid numeric CNPJ bytes without formatting", []byte("11222333000181"), nil},
		{"invalid CNPJ bytes with less than 14 characters", []byte("12.ABC.345/01DE-3"), ErrCNPJMustHave14Characters},
		{"invalid CNPJ bytes with invalid checksum", []byte("12ABC34501DE34"), ErrCNPJNotValid},
		{"invalid numeric CNPJ bytes with invalid checksum", []byte("12.345.678/0001-91"), ErrCNPJNotValid},
		{"invalid CNPJ bytes nil value", nil, ErrCNPJMustHave14Characters},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsCNPJBytes(tt.cnpj)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsCNPJBytes() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalculateCNPJCheckDigits(t *testing.T) {
	tests := []struct {
		name    string
		base    string
		want    string
		wantErr bool
	}{
		{"official alphanumeric example", "12ABC34501DE", "35", false},
		{"official alphanumeric example with formatting", "12.ABC.345/01DE", "35", false},
		{"numeric base with formatting", "04.252.011/0001", "10", false},
		{"numeric base without formatting", "112223330001", "81", false},
		{"numeric base from legacy fixture", "725012630001", "95", false},
		{"numeric base with first check digit as zero", "000000000006", "04", false},
		{"invalid base empty", "", "", true},
		{"invalid base with only zeroes", "000000000000", "", true},
		{"invalid base with lowercase letters", "12abc34501DE", "", true},
		{"invalid base with check digits included", "12ABC34501DE35", "", true},
		{"invalid base with unsupported character", "12ABC34501D@", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateCNPJCheckDigits(tt.base)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CalculateCNPJCheckDigits() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				if !errors.Is(err, ErrCNPJBaseNotValid) {
					t.Fatalf("CalculateCNPJCheckDigits() error = %v, want %v", err, ErrCNPJBaseNotValid)
				}
				return
			}

			if got != tt.want {
				t.Errorf("CalculateCNPJCheckDigits() = %q, want %q", got, tt.want)
			}
		})
	}
}
