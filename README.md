# Multiform Validator Go

Go library made for validating several form fields and common values, such as email, telephone, password, CPF, CNPJ, credit card, real image bytes, and much more.

This package is the Go version of Multiform Validator. It currently includes CPF, CNPJ, email, image, credit card, telephone, empty, blank, ASCII, Base64, CEP, MD5, and port validation, with more validators being added over time. The CNPJ validator supports both the old LEGACY numeric CNPJ format and the new alphanumeric CNPJ format, in accordance with the official Receita Federal / SERPRO specification.

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
)

func main() {
	if err := mv.IsCPF("123.456.789-09"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsCNPJ("12.ABC.345/01DE-35"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsCNPJ("04.252.011/0001-10"); err != nil {
		fmt.Println(err)
	}

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

	fmt.Println(mv.GetOnlyEmail("Contact team: joao@empresa.com, maria@empresa.com"))

	if err := mv.IsCreditCard("4111 1111 1111 1111"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsTelephone("+55 11 91234-5678", "BR"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsEmpty(""); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsBlank("   "); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsCEP("12345-678"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsPort("8080"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsPortNumber(8080); err != nil {
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
	"github.com/Multiform-Validator/go/md5"
	"github.com/Multiform-Validator/go/port"
	"github.com/Multiform-Validator/go/telephone"
	"github.com/Multiform-Validator/go/text"
)
```

## Available Validators

- `IsCPF`
- `IsCNPJ`
- `IsEmail`
- `IsImage`
- `GetOnlyEmail`
- `GetOnlyEmails`
- `IsCreditCard`
- `IsTelephone`
- `IsEmpty`
- `IsBlank`
- `IsAscii`
- `IsBase64`
- `IsCEP`
- `IsMD5`
- `IsPort`
- `IsPortNumber`
- `CalculateCNPJCheckDigits`

`IsTelephone` accepts an optional country argument. Current country-specific validation supports Brazil, United States, China, Japan, Germany, India, United Kingdom, France, Italy, Canada, and South Korea.

## Planned Validators

Some validations are still missing from the Go package documentation and will be implemented soon:

- Password

## Development

```bash
make check
```
