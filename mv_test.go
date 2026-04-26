package mv

import "testing"

func TestIsCPFFromRootPackage(t *testing.T) {
	if err := IsCPF("123.456.789-09"); err != nil {
		t.Fatalf("IsCPF() returned error for valid CPF: %v", err)
	}

	if err := IsCPF("123.456.789-02"); err == nil {
		t.Fatal("IsCPF() expected error for invalid CPF, got nil")
	}
}

func TestIsCPFBytesFromRootPackage(t *testing.T) {
	if err := IsCPFBytes([]byte("123.456.789-09")); err != nil {
		t.Fatalf("IsCPFBytes() returned error for valid CPF: %v", err)
	}

	if err := IsCPFBytes([]byte("123.456.789-02")); err == nil {
		t.Fatal("IsCPFBytes() expected error for invalid CPF, got nil")
	}
}

func TestIsCNPJFromRootPackage(t *testing.T) {
	if err := IsCNPJ("12.ABC.345/01DE-35"); err != nil {
		t.Fatalf("IsCNPJ() returned error for valid CNPJ: %v", err)
	}

	if err := IsCNPJ("12.ABC.345/01DE-34"); err == nil {
		t.Fatal("IsCNPJ() expected error for invalid CNPJ, got nil")
	}
}

func TestIsCNPJBytesFromRootPackage(t *testing.T) {
	if err := IsCNPJBytes([]byte("12.ABC.345/01DE-35")); err != nil {
		t.Fatalf("IsCNPJBytes() returned error for valid CNPJ: %v", err)
	}

	if err := IsCNPJBytes([]byte("12.ABC.345/01DE-34")); err == nil {
		t.Fatal("IsCNPJBytes() expected error for invalid CNPJ, got nil")
	}
}

func TestIsEmailFromRootPackage(t *testing.T) {
	if err := IsEmail("user@example.com"); err != nil {
		t.Fatalf("IsEmail() returned error for valid email: %v", err)
	}

	if err := IsEmail("user.example.com"); err == nil {
		t.Fatal("IsEmail() expected error for invalid email, got nil")
	}
}

func TestIsEmailBytesFromRootPackage(t *testing.T) {
	if err := IsEmailBytes([]byte("user@example.com")); err != nil {
		t.Fatalf("IsEmailBytes() returned error for valid email: %v", err)
	}

	if err := IsEmailBytes([]byte("user.example.com")); err == nil {
		t.Fatal("IsEmailBytes() expected error for invalid email, got nil")
	}
}

func TestIsCreditCardFromRootPackage(t *testing.T) {
	if err := IsCreditCard("4111 1111 1111 1111"); err != nil {
		t.Fatalf("IsCreditCard() returned error for valid credit card: %v", err)
	}

	if err := IsCreditCard("4111 1111 1111 1112"); err == nil {
		t.Fatal("IsCreditCard() expected error for invalid credit card, got nil")
	}
}

func TestIsCreditCardBytesFromRootPackage(t *testing.T) {
	if err := IsCreditCardBytes([]byte("4111 1111 1111 1111")); err != nil {
		t.Fatalf("IsCreditCardBytes() returned error for valid credit card: %v", err)
	}

	if err := IsCreditCardBytes([]byte("4111 1111 1111 1112")); err == nil {
		t.Fatal("IsCreditCardBytes() expected error for invalid credit card, got nil")
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
