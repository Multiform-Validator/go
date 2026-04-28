package mv

import (
	"bytes"
	stdimage "image"
	"image/color"
	"image/png"
	"reflect"
	"testing"

	"github.com/Multiform-Validator/go/email"
)

func TestIsCPFFromRootPackage(t *testing.T) {
	if err := IsCPF("123.456.789-09"); err != nil {
		t.Fatalf("IsCPF() returned error for valid CPF: %v", err)
	}

	if err := IsCPF("123.456.789-02"); err == nil {
		t.Fatal("IsCPF() expected error for invalid CPF, got nil")
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

func TestIsEmailFromRootPackage(t *testing.T) {
	if err := IsEmail("user@example.com"); err != nil {
		t.Fatalf("IsEmail() returned error for valid email: %v", err)
	}

	if err := IsEmail("user.example.com"); err == nil {
		t.Fatal("IsEmail() expected error for invalid email, got nil")
	}
}

func TestIsImageFromRootPackage(t *testing.T) {
	if err := IsImage(validPNGBytes(t)); err != nil {
		t.Fatalf("IsImage() returned error for valid image: %v", err)
	}

	if err := IsImage([]byte("%PDF-1.7")); err == nil {
		t.Fatal("IsImage() expected error for invalid image, got nil")
	}
}

func validPNGBytes(t *testing.T) []byte {
	t.Helper()

	img := stdimage.NewRGBA(stdimage.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{A: 255})

	var buffer bytes.Buffer
	if err := png.Encode(&buffer, img); err != nil {
		t.Fatalf("png.Encode() error = %v", err)
	}

	return buffer.Bytes()
}

func TestGetOnlyEmailFromRootPackage(t *testing.T) {
	got := GetOnlyEmail("Contact team: john@gmail.com, alexa@gmail.com")
	if got != "john@gmail.com" {
		t.Fatalf("GetOnlyEmail() = %q, want %q", got, "john@gmail.com")
	}

	got = GetOnlyEmail("Contact team")
	if got != email.NoEmailFound {
		t.Fatalf("GetOnlyEmail() = %q, want %q", got, email.NoEmailFound)
	}
}

func TestGetOnlyEmailsFromRootPackage(t *testing.T) {
	got := GetOnlyEmails(
		"Contact team: john@gmail.comXTRA, alexa@gmail.comXTRA",
		email.GetOnlyEmailOptions{CleanDomain: true, RepeatEmail: true},
	)
	want := []string{"john@gmail.com", "alexa@gmail.com"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetOnlyEmails() = %#v, want %#v", got, want)
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

func TestIdentifyFlagCardFromRootPackage(t *testing.T) {
	if got := IdentifyFlagCard("4111 1111 1111 1111"); got != "Visa" {
		t.Fatalf("IdentifyFlagCard() = %q, want %q", got, "Visa")
	}

	if got := IdentifyFlagCard("7000000000000000"); got != "Unknown" {
		t.Fatalf("IdentifyFlagCard() = %q, want %q", got, "Unknown")
	}
}

func TestIsTelephoneFromRootPackage(t *testing.T) {
	if err := IsTelephone("+55 11 91234-5678", "BR"); err != nil {
		t.Fatalf("IsTelephone() returned error for valid telephone: %v", err)
	}

	if err := IsTelephone("+91 51234 56789", "IN"); err == nil {
		t.Fatal("IsTelephone() expected error for invalid telephone, got nil")
	}
}

func TestIsEmptyFromRootPackage(t *testing.T) {
	if err := IsEmpty(""); err != nil {
		t.Fatalf("IsEmpty() returned error for empty value: %v", err)
	}

	if err := IsEmpty(" "); err == nil {
		t.Fatal("IsEmpty() expected error for non-empty value, got nil")
	}
}

func TestIsEmptyBytesFromRootPackage(t *testing.T) {
	if err := IsEmptyBytes(nil); err != nil {
		t.Fatalf("IsEmptyBytes() returned error for nil bytes: %v", err)
	}

	if err := IsEmptyBytes([]byte(" ")); err == nil {
		t.Fatal("IsEmptyBytes() expected error for non-empty bytes, got nil")
	}
}

func TestIsBlankFromRootPackage(t *testing.T) {
	if err := IsBlank("   "); err != nil {
		t.Fatalf("IsBlank() returned error for blank value: %v", err)
	}

	if err := IsBlank("value"); err == nil {
		t.Fatal("IsBlank() expected error for non-blank value, got nil")
	}
}

func TestIsBlankBytesFromRootPackage(t *testing.T) {
	if err := IsBlankBytes([]byte("   ")); err != nil {
		t.Fatalf("IsBlankBytes() returned error for blank bytes: %v", err)
	}

	if err := IsBlankBytes([]byte("value")); err == nil {
		t.Fatal("IsBlankBytes() expected error for non-blank bytes, got nil")
	}
}

func TestIsAsciiFromRootPackage(t *testing.T) {
	if err := IsAscii("Hello 123!"); err != nil {
		t.Fatalf("IsAscii() returned error for valid ASCII: %v", err)
	}

	if err := IsAscii("olá"); err == nil {
		t.Fatal("IsAscii() expected error for invalid ASCII, got nil")
	}
}

func TestIsAsciiBytesFromRootPackage(t *testing.T) {
	if err := IsAsciiBytes([]byte("Hello 123!")); err != nil {
		t.Fatalf("IsAsciiBytes() returned error for valid ASCII bytes: %v", err)
	}

	if err := IsAsciiBytes([]byte{0x48, 0x80}); err == nil {
		t.Fatal("IsAsciiBytes() expected error for invalid ASCII bytes, got nil")
	}
}

func TestIsBase64FromRootPackage(t *testing.T) {
	if err := IsBase64("SGVsbG8="); err != nil {
		t.Fatalf("IsBase64() returned error for valid base64: %v", err)
	}

	if err := IsBase64("not base64!"); err == nil {
		t.Fatal("IsBase64() expected error for invalid base64, got nil")
	}
}

func TestIsMACAddressFromRootPackage(t *testing.T) {
	if err := IsMACAddress("00:1A:2B:3C:4D:5E"); err != nil {
		t.Fatalf("IsMACAddress() returned error for valid MAC address: %v", err)
	}

	if err := IsMACAddress("00:1A:2B:3C:4D:ZZ"); err == nil {
		t.Fatal("IsMACAddress() expected error for invalid MAC address, got nil")
	}
}

func TestIsCEPFromRootPackage(t *testing.T) {
	if err := IsCEP("12345-678"); err != nil {
		t.Fatalf("IsCEP() returned error for valid CEP: %v", err)
	}

	if err := IsCEP("12345-67A"); err == nil {
		t.Fatal("IsCEP() expected error for invalid CEP, got nil")
	}
}

func TestIsPostalCodeFromRootPackage(t *testing.T) {
	if err := IsPostalCode("10045-123", "BR"); err != nil {
		t.Fatalf("IsPostalCode() returned error for valid postal code: %v", err)
	}

	if err := IsPostalCode("10045-12", "BR"); err == nil {
		t.Fatal("IsPostalCode() expected error for invalid postal code, got nil")
	}
}

func TestIsMD5FromRootPackage(t *testing.T) {
	if err := IsMD5("d41d8cd98f00b204e9800998ecf8427e"); err != nil {
		t.Fatalf("IsMD5() returned error for valid MD5: %v", err)
	}

	if err := IsMD5("d41d8cd98f00b204e9800998ecf8427g"); err == nil {
		t.Fatal("IsMD5() expected error for invalid MD5, got nil")
	}
}

func TestIsPortFromRootPackage(t *testing.T) {
	if err := IsPort("8080"); err != nil {
		t.Fatalf("IsPort() returned error for valid port: %v", err)
	}

	if err := IsPort("65536"); err == nil {
		t.Fatal("IsPort() expected error for invalid port, got nil")
	}
}

func TestIsPortNumberFromRootPackage(t *testing.T) {
	if err := IsPortNumber(8080); err != nil {
		t.Fatalf("IsPortNumber() returned error for valid port number: %v", err)
	}

	if err := IsPortNumber(65536); err == nil {
		t.Fatal("IsPortNumber() expected error for invalid port number, got nil")
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
