package cnpj

import (
	"errors"
	"fmt"
)

const (
	cnpjSize          = 14
	cnpjSizeWithoutDV = 12
	baseValue         = int('0')
)

var (
	ErrCNPJMustHave14Characters = errors.New("CNPJ must have 14 characters")
	ErrCNPJNotValid             = errors.New("CNPJ is not valid")
	ErrCNPJBaseNotValid         = errors.New("CNPJ is not valid for check digit calculation")
)

var cnpjCheckDigitWeights = [13]int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

func IsCNPJ(cnpj string) error {
	digits, count := extractCNPJCharacters(cnpj, cnpjSize)
	if count != cnpjSize {
		return ErrCNPJMustHave14Characters
	}

	if !isCNPJFormationValidWithDVBytes(digits) {
		return ErrCNPJNotValid
	}

	first := calculateDigitBytes(digits[:cnpjSizeWithoutDV])
	second := calculateDigitBytesWithExtra(digits[:cnpjSizeWithoutDV], first)
	if digits[cnpjSizeWithoutDV] != digitToByte(first) || digits[cnpjSizeWithoutDV+1] != digitToByte(second) {
		return ErrCNPJNotValid
	}

	return nil
}

func CalculateCNPJCheckDigits(baseCNPJ string) (string, error) {
	baseCNPJ = removeFormattingCharacters(baseCNPJ)
	if !isCNPJFormationValidWithoutDV(baseCNPJ) {
		return "", fmt.Errorf("%w: %s", ErrCNPJBaseNotValid, baseCNPJ)
	}

	first := calculateDigit(baseCNPJ)
	second := calculateDigitWithExtra(baseCNPJ, first)

	return string([]byte{digitToByte(first), digitToByte(second)}), nil
}

func removeFormattingCharacters(cnpj string) string {
	cnpj = trim(cnpj)
	for i := 0; i < len(cnpj); i++ {
		switch cnpj[i] {
		case '.', '/', '-':
			cleaned := make([]byte, 0, len(cnpj)-1)
			cleaned = append(cleaned, cnpj[:i]...)
			for ; i < len(cnpj); i++ {
				switch cnpj[i] {
				case '.', '/', '-':
					continue
				default:
					cleaned = append(cleaned, cnpj[i])
				}
			}
			return string(cleaned)
		}
	}

	return cnpj
}

func extractCNPJCharacters(cnpj string, max int) ([cnpjSize]byte, int) {
	var digits [cnpjSize]byte
	cnpj = trim(cnpj)
	count := 0
	for i := 0; i < len(cnpj); i++ {
		switch cnpj[i] {
		case '.', '/', '-':
			continue
		}

		if count == max {
			return digits, count + 1
		}

		digits[count] = cnpj[i]
		count++
	}

	return digits, count
}

func trim(value string) string {
	start := 0
	for start < len(value) && value[start] <= ' ' {
		start++
	}

	end := len(value)
	for end > start && value[end-1] <= ' ' {
		end--
	}

	return value[start:end]
}

func isCNPJFormationValidWithoutDV(cnpj string) bool {
	return len(cnpj) == cnpjSizeWithoutDV &&
		hasOnlyUppercaseLettersAndDigits(cnpj) &&
		!hasOnlyZeroes(cnpj)
}

func isCNPJFormationValidWithDVBytes(cnpj [cnpjSize]byte) bool {
	return hasValidBaseFormationBytes(cnpj) &&
		isDigit(cnpj[cnpjSizeWithoutDV]) &&
		isDigit(cnpj[cnpjSizeWithoutDV+1]) &&
		!hasOnlyZeroesBytes(cnpj)
}

func hasValidBaseFormationBytes(cnpj [cnpjSize]byte) bool {
	for i := 0; i < cnpjSizeWithoutDV; i++ {
		if !isUppercaseLetterOrDigit(cnpj[i]) {
			return false
		}
	}

	return true
}

func hasOnlyUppercaseLettersAndDigits(value string) bool {
	for i := 0; i < len(value); i++ {
		c := value[i]
		if (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			continue
		}

		return false
	}

	return true
}

func hasOnlyZeroes(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] != '0' {
			return false
		}
	}

	return true
}

func hasOnlyZeroesBytes(value [cnpjSize]byte) bool {
	for i := 0; i < len(value); i++ {
		if value[i] != '0' {
			return false
		}
	}

	return true
}

func isUppercaseLetterOrDigit(value byte) bool {
	return (value >= 'A' && value <= 'Z') || isDigit(value)
}

func isDigit(value byte) bool {
	return value >= '0' && value <= '9'
}

func digitToByte(value int) byte {
	return "0123456789"[value]
}

func calculateDigit(cnpj string) int {
	sum := 0
	for index := len(cnpj) - 1; index >= 0; index-- {
		characterValue := int(cnpj[index]) - baseValue
		sum += characterValue * cnpjCheckDigitWeights[len(cnpjCheckDigitWeights)-len(cnpj)+index]
	}

	if sum%11 < 2 {
		return 0
	}

	return 11 - (sum % 11)
}

func calculateDigitBytes(cnpj []byte) int {
	sum := 0
	for index := len(cnpj) - 1; index >= 0; index-- {
		characterValue := int(cnpj[index]) - baseValue
		sum += characterValue * cnpjCheckDigitWeights[len(cnpjCheckDigitWeights)-len(cnpj)+index]
	}

	if sum%11 < 2 {
		return 0
	}

	return 11 - (sum % 11)
}

func calculateDigitBytesWithExtra(cnpj []byte, extraDigit int) int {
	sum := 0
	totalLength := len(cnpj) + 1
	for index := len(cnpj) - 1; index >= 0; index-- {
		characterValue := int(cnpj[index]) - baseValue
		sum += characterValue * cnpjCheckDigitWeights[len(cnpjCheckDigitWeights)-totalLength+index]
	}
	sum += extraDigit * cnpjCheckDigitWeights[len(cnpjCheckDigitWeights)-1]

	if sum%11 < 2 {
		return 0
	}

	return 11 - (sum % 11)
}

func calculateDigitWithExtra(cnpj string, extraDigit int) int {
	sum := 0
	totalLength := len(cnpj) + 1
	for index := len(cnpj) - 1; index >= 0; index-- {
		characterValue := int(cnpj[index]) - baseValue
		sum += characterValue * cnpjCheckDigitWeights[len(cnpjCheckDigitWeights)-totalLength+index]
	}
	sum += extraDigit * cnpjCheckDigitWeights[len(cnpjCheckDigitWeights)-1]

	if sum%11 < 2 {
		return 0
	}

	return 11 - (sum % 11)
}
