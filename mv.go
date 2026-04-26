package mv

import (
	"github.com/Multiform-Validator/go/cnpj"
	"github.com/Multiform-Validator/go/cpf"
	"github.com/Multiform-Validator/go/creditcard"
	"github.com/Multiform-Validator/go/email"
	"github.com/Multiform-Validator/go/text"
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

func IsEmail(value string) error {
	return email.IsEmail(value)
}

func IsEmailBytes(value []byte) error {
	return email.IsEmailBytes(value)
}

func IsCreditCard(value string) error {
	return creditcard.IsCreditCard(value)
}

func IsCreditCardBytes(value []byte) error {
	return creditcard.IsCreditCardBytes(value)
}

func IsEmpty(value string) error {
	return text.IsEmpty(value)
}

func IsEmptyBytes(value []byte) error {
	return text.IsEmptyBytes(value)
}

func IsBlank(value string) error {
	return text.IsBlank(value)
}

func IsBlankBytes(value []byte) error {
	return text.IsBlankBytes(value)
}

func CalculateCNPJCheckDigits(value string) (string, error) {
	return cnpj.CalculateCNPJCheckDigits(value)
}
