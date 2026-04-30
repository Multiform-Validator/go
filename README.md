# Multiform Validator Go

Go library made for validating several form fields and common values, such as ASCII, Base64, CEP, CNPJ, CPF, credit card, email, real image bytes, MAC address, MD5, port, postal code, telephone, password, and much more.

This package is the Go version of Multiform Validator. It currently includes ASCII, Base64, CEP, CNPJ, CPF, credit card, email, image, MAC address, MD5, port, postal code, telephone, text, and higher-level validation helpers, with more validators being added over time. The CNPJ validator supports both the old LEGACY numeric CNPJ format and the new alphanumeric CNPJ format, in accordance with the official Receita Federal / SERPRO specification.

## Install

```bash
go get github.com/Multiform-Validator/go
```

## Usage

```go
package main

import (
	"fmt"
	"os"

	mv "github.com/Multiform-Validator/go"
	"github.com/Multiform-Validator/go/validate"
)

func main() {
	if err := mv.IsCEP("12345-678"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsCNPJ("12.ABC.345/01DE-35"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsCNPJ("04.252.011/0001-10"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsCPF("123.456.789-09"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(mv.IdentifyFlagCard("4111 1111 1111 1111"))

	if err := mv.IsCreditCard("4111 1111 1111 1111"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(mv.GetOnlyEmail("Contact team: joao@empresa.com, maria@empresa.com"))

	if err := mv.IsEmail("user@example.com"); err != nil {
		fmt.Println(err)
	}

	imageBytes, err := os.ReadFile("avatar.png")
	if err != nil {
		fmt.Println(err)
	}

	if err := mv.IsImage(imageBytes); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsMACAddress("00:1A:2B:3C:4D:5E"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsPort("8080"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsPortNumber(8080); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsPostalCode("10045-123", "BR"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsTelephone("+55 11 91234-5678", "BR"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsBlank("   "); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsEmpty(""); err != nil {
		fmt.Println(err)
	}

	if err := validate.Email("user@gmail.com", validate.EmailOptions{ValidDomains: true}); err != nil {
		fmt.Println(err)
	}

	if err := validate.Password("MyP@ssw0rd", validate.PasswordOptions{
		MinLength:          8,
		MaxLength:          20,
		RequireUppercase:   true,
		RequireSpecialChar: true,
		RequireNumber:      true,
		RequireLetter:      true,
	}); err != nil {
		fmt.Println(err)
	}
}
```

You can also import each validator package directly:

```go
import (
	"github.com/Multiform-Validator/go/ascii"
	"github.com/Multiform-Validator/go/base64"
	"github.com/Multiform-Validator/go/cep"
	"github.com/Multiform-Validator/go/cnpj"
	"github.com/Multiform-Validator/go/cpf"
	"github.com/Multiform-Validator/go/creditcard"
	"github.com/Multiform-Validator/go/email"
	"github.com/Multiform-Validator/go/image"
	"github.com/Multiform-Validator/go/macaddress"
	"github.com/Multiform-Validator/go/md5"
	"github.com/Multiform-Validator/go/port"
	"github.com/Multiform-Validator/go/postalcode"
	"github.com/Multiform-Validator/go/telephone"
	"github.com/Multiform-Validator/go/text"
	"github.com/Multiform-Validator/go/validate"
)
```

## Validator Packages

Most validators follow the direct `Is*` style, such as `IsEmail`, `IsCPF`, and `IsPort`. These are simple checks for one specific value.

The `validate` package is intentionally separate and works a little differently. It is for higher-level validators with options and composed rules, using the `validate.*` style, such as `validate.Email` and `validate.Password`.

## Available Validators

- `IsAscii`
- `IsAsciiBytes`
- `IsBase64`
- `IsCEP`
- `CalculateCNPJCheckDigits`
- `IsCNPJ`
- `IsCPF`
- `IdentifyFlagCard`
- `IsCreditCard`
- `GetOnlyEmail`
- `GetOnlyEmails`
- `IsEmail`
- `IsImage`
- `IsMACAddress`
- `IsMD5`
- `IsPort`
- `IsPortNumber`
- `IsPostalCode`
- `IsTelephone`
- `IsBlank`
- `IsBlankBytes`
- `IsEmpty`
- `IsEmptyBytes`
- `validate.Email`
- `validate.Password`

`IsTelephone` accepts an optional country argument. Current country-specific validation supports Brazil, United States, China, Japan, Germany, India, United Kingdom, France, Italy, Canada, and South Korea.

`IsPostalCode` accepts an optional country argument. Current country-specific validation supports Brazil, United States, Canada, United Kingdom, France, Netherlands, Japan, Spain, South Africa, Germany, Switzerland, and Italy.

`IdentifyFlagCard` is a special helper from the credit card package, similar to how `GetOnlyEmail` belongs to the email package. It identifies the card flag by prefix and returns `Unknown` when no known flag matches.

`validate.Email` provides higher-level email validation options inspired by the TypeScript package: max length, country suffix, default allowed domains, and custom allowed domains.

`validate.Password` validates password length and optional uppercase, special character, number, and letter requirements.

## Development

```bash
make check
```
