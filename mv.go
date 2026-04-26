package mv

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

func GetOnlyEmail(value string, options ...email.GetOnlyEmailOptions) string {
	return email.GetOnlyEmail(value, options...)
}

func GetOnlyEmails(value string, options ...email.GetOnlyEmailOptions) []string {
	return email.GetOnlyEmails(value, options...)
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

func IsAscii(value string) error {
	return ascii.IsAscii(value)
}

func IsAsciiBytes(value []byte) error {
	return ascii.IsAsciiBytes(value)
}

func IsBase64(value string) error {
	return base64.IsBase64(value)
}

func IsBase64Bytes(value []byte) error {
	return base64.IsBase64Bytes(value)
}

func IsCEP(value string) error {
	return cep.IsCEP(value)
}

func IsCEPBytes(value []byte) error {
	return cep.IsCEPBytes(value)
}

func IsMD5(value string) error {
	return md5.IsMD5(value)
}

func IsMD5Bytes(value []byte) error {
	return md5.IsMD5Bytes(value)
}

func IsPort(value string) error {
	return port.IsPort(value)
}

func IsPortBytes(value []byte) error {
	return port.IsPortBytes(value)
}

func CalculateCNPJCheckDigits(value string) (string, error) {
	return cnpj.CalculateCNPJCheckDigits(value)
}
