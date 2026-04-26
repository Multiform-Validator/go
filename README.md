# Multiform Validator Go

Go library made for validating several form fields and common values, such as email, telephone, password, CPF, CNPJ, credit card, image MIME type magic numbers, and much more.

This package is the Go version of Multiform Validator. It currently includes CPF and CNPJ validation, with more validators being added over time. The CNPJ validator supports both the classic numeric format and the new alphanumeric format.

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
	if err := mv.IsCPFValid("123.456.789-09"); err != nil {
		fmt.Println(err)
	}

	if err := mv.IsCNPJValid("12.ABC.345/01DE-35"); err != nil {
		fmt.Println(err)
	}
}
```

You can also import each validator package directly:

```go
import (
	"github.com/Multiform-Validator/go/cnpj"
	"github.com/Multiform-Validator/go/cpf"
)
```

## Available Validators

- `IsCPFValid`
- `IsCPFValidBytes`
- `IsCNPJValid`
- `IsCNPJValidBytes`
- `CalculateCNPJCheckDigits`

## Development

```bash
make check
```
