package cep

import (
	"errors"
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
	start := 0
	for start < len(cep) && cep[start] <= ' ' {
		start++
	}

	end := len(cep)
	for end > start && cep[end-1] <= ' ' {
		end--
	}

	count := 0
	onlyDigits := true
	for i := start; i < end; i++ {
		switch cep[i] {
		case '-':
			continue
		case '.', '/', ' ', '\t', '\n', '\r':
			return 0, false, false
		}

		if cep[i] < '0' || cep[i] > '9' {
			onlyDigits = false
		}
		count++
	}

	return count, onlyDigits, true
}
