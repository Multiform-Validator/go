package mv

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

func IsCPF(value string) error {
	return cpf.IsCPF(value)
}

func IsCNPJ(value string) error {
	return cnpj.IsCNPJ(value)
}

func IsEmail(value string) error {
	return email.IsEmail(value)
}

func IsImage(value []byte) error {
	return image.IsImage(value)
}

func GetOnlyEmail(value string, options ...email.GetOnlyEmailOptions) string {
	return email.GetOnlyEmail(value, options...)
}

func GetOnlyEmails(value string, options ...email.GetOnlyEmailOptions) []string {
	return email.GetOnlyEmails(value, options...)
}

func IsCreditCard(value string) error {
	return creditcard.IsCreditCard(value)
}

func IsTelephone(value string, countries ...string) error {
	return telephone.IsTelephone(value, countries...)
}

func IsEmpty(value string) error {
	return text.IsEmpty(value)
}

func IsBlank(value string) error {
	return text.IsBlank(value)
}

func IsAscii(value string) error {
	return ascii.IsAscii(value)
}

func IsBase64(value string) error {
	return base64.IsBase64(value)
}

func IsCEP(value string) error {
	return cep.IsCEP(value)
}

func IsMD5(value string) error {
	return md5.IsMD5(value)
}

func IsPort(value string) error {
	return port.IsPort(value)
}

func IsPortNumber(value int) error {
	return port.IsPortNumber(value)
}

func CalculateCNPJCheckDigits(value string) (string, error) {
	return cnpj.CalculateCNPJCheckDigits(value)
}
