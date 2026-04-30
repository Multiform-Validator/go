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
	count, onlyDigits, ok := analyzeCEPCharacters(cep)
	if !ok {
		return ErrCEPNotValid
	}

	if count != cepSize {
		return ErrCEPMustHave8Digits
	}

	if !onlyDigits {
		return ErrCEPNotValid
	}

	return nil
}

func analyzeCEPCharacters(cep string) (int, bool, bool) {
	cep = strings.TrimSpace(cep)
	if strings.ContainsAny(cep, "./ ") {
		return 0, false, false
	}

	count := 0
	onlyDigits := true
	for i := 0; i < len(cep); i++ {
		if cep[i] == '-' {
			continue
		}
		if cep[i] < '0' || cep[i] > '9' {
			onlyDigits = false
		}
		count++
	}

	return count, onlyDigits, true
}
