package mv_test

import (
	"fmt"

	mv "github.com/Multiform-Validator/go"
	"github.com/Multiform-Validator/go/email"
	"github.com/Multiform-Validator/go/validate"
)

func ExampleIsCPF() {
	fmt.Println(mv.IsCPF("123.456.789-09") == nil)
	fmt.Println(mv.IsCPF("123.456.789-02") == nil)

	// Output:
	// true
	// false
}

func ExampleIsCNPJ() {
	fmt.Println(mv.IsCNPJ("04.252.011/0001-10") == nil)
	fmt.Println(mv.IsCNPJ("12.345.678/0001-91") == nil)

	// Output:
	// true
	// false
}

func ExampleIsEmail() {
	fmt.Println(mv.IsEmail("user@example.com") == nil)
	fmt.Println(mv.IsEmail("user.example.com") == nil)

	// Output:
	// true
	// false
}

func ExampleGetOnlyEmail() {
	fmt.Println(mv.GetOnlyEmail("Contact team: joao@empresa.com, maria@empresa.com"))

	// Output:
	// joao@empresa.com
}

func ExampleGetOnlyEmails() {
	emails := mv.GetOnlyEmails(
		"Contact: john@gmail.comXTRA, alexa@gmail.comXTRA",
		email.GetOnlyEmailOptions{CleanDomain: true, RepeatEmail: true},
	)
	fmt.Println(emails)

	// Output:
	// [john@gmail.com alexa@gmail.com]
}

func ExampleIsCreditCard() {
	fmt.Println(mv.IsCreditCard("4111 1111 1111 1111") == nil)
	fmt.Println(mv.IsCreditCard("4111 1111 1111 1112") == nil)

	// Output:
	// true
	// false
}

func ExampleIdentifyFlagCard() {
	fmt.Println(mv.IdentifyFlagCard("4111 1111 1111 1111"))
	fmt.Println(mv.IdentifyFlagCard("7000 0000 0000 0000"))

	// Output:
	// Visa
	// Unknown
}

func ExampleIsTelephone() {
	fmt.Println(mv.IsTelephone("+55 11 91234-5678", "BR") == nil)
	fmt.Println(mv.IsTelephone("+91 51234 56789", "IN") == nil)

	// Output:
	// true
	// false
}

func ExampleIsBase64() {
	fmt.Println(mv.IsBase64("SGVsbG8gV29ybGQ=") == nil)
	fmt.Println(mv.IsBase64("SGVs bG8=") == nil)

	// Output:
	// true
	// false
}

func ExampleIsMD5() {
	fmt.Println(mv.IsMD5("d41d8cd98f00b204e9800998ecf8427e") == nil)
	fmt.Println(mv.IsMD5("d41d8cd98f00b204e9800998ecf8427g") == nil)

	// Output:
	// true
	// false
}

func ExampleIsCEP() {
	fmt.Println(mv.IsCEP("12345-678") == nil)
	fmt.Println(mv.IsCEP("12345.678") == nil)

	// Output:
	// true
	// false
}

func ExampleIsPostalCode() {
	fmt.Println(mv.IsPostalCode("SW1A 1AA", "UK") == nil)
	fmt.Println(mv.IsPostalCode("90210-123", "US") == nil)

	// Output:
	// true
	// false
}

func ExampleIsPort() {
	fmt.Println(mv.IsPort("8080") == nil)
	fmt.Println(mv.IsPort("65536") == nil)

	// Output:
	// true
	// false
}

func ExampleIsBlank() {
	fmt.Println(mv.IsBlank(" \t\n") == nil)
	fmt.Println(mv.IsBlank("value") == nil)

	// Output:
	// true
	// false
}

func ExampleIsAscii() {
	fmt.Println(mv.IsAscii("Hello 123!") == nil)
	fmt.Println(mv.IsAscii("olá") == nil)

	// Output:
	// true
	// false
}

func ExampleCalculateCNPJCheckDigits() {
	digits, err := mv.CalculateCNPJCheckDigits("12ABC34501DE")
	fmt.Println(digits, err == nil)

	// Output:
	// 35 true
}

func Example_validateEmail() {
	err := validate.Email("user@gmail.com", validate.EmailOptions{ValidDomains: true})
	fmt.Println(err == nil)

	// Output:
	// true
}

func Example_validatePassword() {
	err := validate.Password("MyP@ssw0rd", validate.PasswordOptions{
		MinLength:          8,
		MaxLength:          20,
		RequireUppercase:   true,
		RequireSpecialChar: true,
		RequireNumber:      true,
		RequireLetter:      true,
	})
	fmt.Println(err == nil)

	// Output:
	// true
}
