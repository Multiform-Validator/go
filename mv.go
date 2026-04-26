package mv

import (
	"github.com/Multiform-Validator/go/cnpj"
	"github.com/Multiform-Validator/go/cpf"
)

func IsCPF(value string) error {
	return cpf.IsCPF(value)
}

func IsCPFBytes(value []byte) error {
	return cpf.IsCPFBytes(value)
}

func IsCNPJ(value string) error {
	return cnpj.IsCNPJ(value)
}

func IsCNPJBytes(value []byte) error {
	return cnpj.IsCNPJBytes(value)
}

func CalculateCNPJCheckDigits(value string) (string, error) {
	return cnpj.CalculateCNPJCheckDigits(value)
}
