# Multiform Validator Go

Go library made for validating several form fields and common values, such as email, telephone, password, CPF, CNPJ, credit card, image MIME type magic numbers, and much more.

This package is the Go version of Multiform Validator. It currently includes CPF, CNPJ, email, credit card, empty, blank, ASCII, Base64, CEP, MD5, and port validation, with more validators being added over time. The CNPJ validator supports both the classic numeric format and the new alphanumeric format.

## Install

```bash
go get github.com/Multiform-Validator/go
```

## Usage

```go
package main

import (
	"fmt"

	mv "github.com/Multiform-Validator/go"
)

func main() {
	if err := mv.IsCPF("123.456.789-09"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsCNPJ("12.ABC.345/01DE-35"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsEmail("user@example.com"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(mv.GetOnlyEmail("Contact team: joao@empresa.com, maria@empresa.com"))

	if err := mv.IsCreditCard("4111 1111 1111 1111"); err != nil {
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
	"github.com/Multiform-Validator/go/md5"
	"github.com/Multiform-Validator/go/port"
	"github.com/Multiform-Validator/go/text"
)
```

## Available Validators

- `IsCPF`
- `IsCPFBytes`
- `IsCNPJ`
- `IsCNPJBytes`
- `IsEmail`
- `IsEmailBytes`
- `GetOnlyEmail`
- `GetOnlyEmails`
- `IsCreditCard`
- `IsCreditCardBytes`
- `IsEmpty`
- `IsEmptyBytes`
- `IsBlank`
- `IsBlankBytes`
- `IsAscii`
- `IsAsciiBytes`
- `IsBase64`
- `IsBase64Bytes`
- `IsCEP`
- `IsCEPBytes`
- `IsMD5`
- `IsMD5Bytes`
- `IsPort`
- `IsPortBytes`
- `CalculateCNPJCheckDigits`

## Development

```bash
make check
```
