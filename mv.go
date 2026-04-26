package mv

import (
	"github.com/Multiform-Validator/go/cnpj"
	"github.com/Multiform-Validator/go/cpf"
)

func IsCPFValid(value string) error {
	return cpf.IsCPFValid(value)
}

func IsCPFValidBytes(value []byte) error {
	return cpf.IsCPFValidBytes(value)
}

func IsCNPJValid(value string) error {
	return cnpj.IsCNPJValid(value)
}

func IsCNPJValidBytes(value []byte) error {
	return cnpj.IsCNPJValidBytes(value)
}

func CalculateCNPJCheckDigits(value string) (string, error) {
	return cnpj.CalculateCNPJCheckDigits(value)
}
