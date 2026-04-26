package cpf

import (
	"errors"
)

var (
	ErrCPFMustHave11Digits = errors.New("CPF must have 11 digits")
	ErrCPFNotValid         = errors.New("CPF is not valid")
)

func IsCPF(cpf string) error {
	digits, count := extractDigitsFromString(cpf)
	if count != 11 {
		return ErrCPFMustHave11Digits
	}

	return validateCPFDigits(digits)
}

func IsCPFBytes(cpf []byte) error {
	digits, count := extractDigitsFromBytes(cpf)
	if count != 11 {
		return ErrCPFMustHave11Digits
	}

	return validateCPFDigits(digits)
}

func extractDigitsFromString(input string) ([11]byte, int) {
	var digits [11]byte
	count := 0
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c < '0' || c > '9' {
			continue
		}

		if count == 11 {
			return digits, count + 1
		}

		digits[count] = c - '0'
		count++
	}

	return digits, count
}

func extractDigitsFromBytes(input []byte) ([11]byte, int) {
	var digits [11]byte
	count := 0
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c < '0' || c > '9' {
			continue
		}

		if count == 11 {
			return digits, count + 1
		}

		digits[count] = c - '0'
		count++
	}

	return digits, count
}

func hasAllRepeatedDigits(digits [11]byte) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i] != digits[0] {
			return false
		}
	}

	return true
}

func calculateCheckDigit(digits [11]byte, size int, weightStart int) int {
	sum := 0
	weight := weightStart
	for i := 0; i < size; i++ {
		sum += int(digits[i]) * weight
		weight--
	}

	remainder := (sum * 10) % 11
	if remainder == 10 {
		return 0
	}

	return remainder
}

func validateCPFDigits(digits [11]byte) error {
	if hasAllRepeatedDigits(digits) {
		return ErrCPFNotValid
	}

	first := calculateCheckDigit(digits, 9, 10)
	if first != int(digits[9]) {
		return ErrCPFNotValid
	}

	second := calculateCheckDigit(digits, 10, 11)
	if second != int(digits[10]) {
		return ErrCPFNotValid
	}

	return nil
}
