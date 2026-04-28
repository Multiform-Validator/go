package mv_test

import (
	"fmt"

	mv "github.com/Multiform-Validator/go"
	"github.com/Multiform-Validator/go/validate"
)

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

func ExampleIdentifyFlagCard() {
	fmt.Println(mv.IdentifyFlagCard("4111 1111 1111 1111"))
	fmt.Println(mv.IdentifyFlagCard("7000 0000 0000 0000"))

	// Output:
	// Visa
	// Unknown
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
