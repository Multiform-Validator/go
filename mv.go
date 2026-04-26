package mv

import "github.com/Multiform-Validator/go/cpf"

func IsCPFValid(value string) error {
	return cpf.IsCPFValid(value)
}

func IsCPFValidBytes(value []byte) error {
	return cpf.IsCPFValidBytes(value)
}
