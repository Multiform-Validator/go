package mv

import "testing"

func TestIsCPFValidFromRootPackage(t *testing.T) {
	if err := IsCPFValid("123.456.789-09"); err != nil {
		t.Fatalf("IsCPFValid() returned error for valid CPF: %v", err)
	}

	if err := IsCPFValid("123.456.789-02"); err == nil {
		t.Fatal("IsCPFValid() expected error for invalid CPF, got nil")
	}
}

func TestIsCPFValidBytesFromRootPackage(t *testing.T) {
	if err := IsCPFValidBytes([]byte("123.456.789-09")); err != nil {
		t.Fatalf("IsCPFValidBytes() returned error for valid CPF: %v", err)
	}

	if err := IsCPFValidBytes([]byte("123.456.789-02")); err == nil {
		t.Fatal("IsCPFValidBytes() expected error for invalid CPF, got nil")
	}
}

func TestIsCNPJValidFromRootPackage(t *testing.T) {
	if err := IsCNPJValid("12.ABC.345/01DE-35"); err != nil {
		t.Fatalf("IsCNPJValid() returned error for valid CNPJ: %v", err)
	}

	if err := IsCNPJValid("12.ABC.345/01DE-34"); err == nil {
		t.Fatal("IsCNPJValid() expected error for invalid CNPJ, got nil")
	}
}

func TestIsCNPJValidBytesFromRootPackage(t *testing.T) {
	if err := IsCNPJValidBytes([]byte("12.ABC.345/01DE-35")); err != nil {
		t.Fatalf("IsCNPJValidBytes() returned error for valid CNPJ: %v", err)
	}

	if err := IsCNPJValidBytes([]byte("12.ABC.345/01DE-34")); err == nil {
		t.Fatal("IsCNPJValidBytes() expected error for invalid CNPJ, got nil")
	}
}

func TestCalculateCNPJCheckDigitsFromRootPackage(t *testing.T) {
	got, err := CalculateCNPJCheckDigits("12ABC34501DE")
	if err != nil {
		t.Fatalf("CalculateCNPJCheckDigits() returned error for valid CNPJ base: %v", err)
	}

	if got != "35" {
		t.Fatalf("CalculateCNPJCheckDigits() = %q, want %q", got, "35")
	}

	if _, err := CalculateCNPJCheckDigits("12ABC34501DE35"); err == nil {
		t.Fatal("CalculateCNPJCheckDigits() expected error for invalid CNPJ base, got nil")
	}
}
