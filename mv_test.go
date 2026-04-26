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
