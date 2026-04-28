package mv_test

import (
	"bytes"
	stdimage "image"
	"image/color"
	"image/png"
	"testing"

	mv "github.com/Multiform-Validator/go"
	"github.com/Multiform-Validator/go/email"
	"github.com/Multiform-Validator/go/validate"
)

var (
	benchErr     error
	benchString  string
	benchStrings []string
)

func BenchmarkIsCPF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsCPF("123.456.789-09")
	}
}

func BenchmarkIsCNPJ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsCNPJ("12.ABC.345/01DE-35")
	}
}

func BenchmarkIsEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsEmail("user@example.com")
	}
}

func BenchmarkIsImage(b *testing.B) {
	value := benchmarkPNGBytes(b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchErr = mv.IsImage(value)
	}
}

func BenchmarkGetOnlyEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchString = mv.GetOnlyEmail("Contact team: john@gmail.com, alexa@gmail.com")
	}
}

func BenchmarkGetOnlyEmails(b *testing.B) {
	options := email.GetOnlyEmailOptions{CleanDomain: true, RepeatEmail: true}
	for i := 0; i < b.N; i++ {
		benchStrings = mv.GetOnlyEmails(
			"Contact team: john@gmail.comXTRA, alexa@gmail.comXTRA",
			options,
		)
	}
}

func BenchmarkIsCreditCard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsCreditCard("4111 1111 1111 1111")
	}
}

func BenchmarkIdentifyFlagCard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchString = mv.IdentifyFlagCard("4111 1111 1111 1111")
	}
}

func BenchmarkIsTelephone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsTelephone("+55 11 91234-5678", "BR")
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsEmpty("")
	}
}

func BenchmarkIsEmptyBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsEmptyBytes(nil)
	}
}

func BenchmarkIsBlank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsBlank(" \t\n")
	}
}

func BenchmarkIsBlankBytes(b *testing.B) {
	value := []byte(" \t\n")
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsBlankBytes(value)
	}
}

func BenchmarkIsAscii(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsAscii("Hello 123!")
	}
}

func BenchmarkIsAsciiBytes(b *testing.B) {
	value := []byte("Hello 123!")
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsAsciiBytes(value)
	}
}

func BenchmarkIsBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsBase64("SGVsbG8gV29ybGQ=")
	}
}

func BenchmarkIsMACAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsMACAddress("00:1A:2B:3C:4D:5E")
	}
}

func BenchmarkIsCEP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsCEP("12345-678")
	}
}

func BenchmarkIsPostalCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsPostalCode("SW1A 1AA", "UK")
	}
}

func BenchmarkIsMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsMD5("d41d8cd98f00b204e9800998ecf8427e")
	}
}

func BenchmarkIsPort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsPort("8080")
	}
}

func BenchmarkIsPortNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchErr = mv.IsPortNumber(8080)
	}
}

func BenchmarkCalculateCNPJCheckDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchString, benchErr = mv.CalculateCNPJCheckDigits("12ABC34501DE")
	}
}

func BenchmarkValidateEmail(b *testing.B) {
	options := validate.EmailOptions{ValidDomains: true}
	for i := 0; i < b.N; i++ {
		benchErr = validate.Email("user@gmail.com", options)
	}
}

func BenchmarkValidatePassword(b *testing.B) {
	options := validate.PasswordOptions{
		MinLength:          8,
		MaxLength:          20,
		RequireUppercase:   true,
		RequireSpecialChar: true,
		RequireNumber:      true,
		RequireLetter:      true,
	}
	for i := 0; i < b.N; i++ {
		benchErr = validate.Password("MyP@ssw0rd", options)
	}
}

func benchmarkPNGBytes(b *testing.B) []byte {
	b.Helper()

	img := stdimage.NewRGBA(stdimage.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{A: 255})

	var buffer bytes.Buffer
	if err := png.Encode(&buffer, img); err != nil {
		b.Fatalf("png.Encode() error = %v", err)
	}

	return buffer.Bytes()
}
