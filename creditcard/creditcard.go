package creditcard

import (
	"errors"
)

const (
	minCreditCardDigits = 12
	maxCreditCardDigits = 19
)

var (
	ErrCreditCardMustHaveBetween12And19Digits = errors.New("credit card must have between 12 and 19 digits")
	ErrCreditCardNotValid                     = errors.New("credit card is not valid")
)

func IsCreditCard(creditCard string) error {
	digits, count, ok := extractDigitsFromString(creditCard)
	if !ok {
		return ErrCreditCardNotValid
	}

	if count < minCreditCardDigits || count > maxCreditCardDigits {
		return ErrCreditCardMustHaveBetween12And19Digits
	}

	return validateCreditCardDigits(digits[:count])
}

func extractDigitsFromString(input string) ([maxCreditCardDigits]byte, int, bool) {
	var digits [maxCreditCardDigits]byte
	count := 0
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c == ' ' || c == '-' {
			continue
		}

		if c < '0' || c > '9' {
			return digits, count, false
		}

		if count == maxCreditCardDigits {
			return digits, count + 1, true
		}

		digits[count] = c - '0'
		count++
	}

	return digits, count, true
}

func validateCreditCardDigits(digits []byte) error {
	if hasAllRepeatedDigits(digits) || !passesLuhn(digits) {
		return ErrCreditCardNotValid
	}

	return nil
}

func hasAllRepeatedDigits(digits []byte) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i] != digits[0] {
			return false
		}
	}

	return true
}

func passesLuhn(digits []byte) bool {
	sum := 0
	shouldDouble := false
	for i := len(digits) - 1; i >= 0; i-- {
		digit := int(digits[i])
		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		shouldDouble = !shouldDouble
	}

	return sum%10 == 0
}
