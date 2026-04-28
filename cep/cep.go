package cep

import (
	"errors"
	"strings"
)

const cepSize = 8

var (
	ErrCEPMustHave8Digits = errors.New("CEP must have 8 digits")
	ErrCEPNotValid        = errors.New("CEP is not valid")
)

func IsCEP(cep string) error {
	cep, ok := removeFormattingCharacters(cep)
	if !ok {
		return ErrCEPNotValid
	}

	if len(cep) != cepSize {
		return ErrCEPMustHave8Digits
	}

	if !hasOnlyDigits(cep) {
		return ErrCEPNotValid
	}

	return nil
}

func removeFormattingCharacters(cep string) (string, bool) {
	cep = strings.TrimSpace(cep)
	if strings.ContainsAny(cep, "./ ") {
		return "", false
	}

	replacer := strings.NewReplacer("-", "")
	return replacer.Replace(cep), true
}

func hasOnlyDigits(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] < '0' || value[i] > '9' {
			return false
		}
	}

	return true
}
