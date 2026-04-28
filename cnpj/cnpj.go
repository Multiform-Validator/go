package cnpj

import (
	"errors"
	"fmt"
	"strings"
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
	cnpj = removeFormattingCharacters(cnpj)
	if len(cnpj) != cnpjSize {
		return ErrCNPJMustHave14Characters
	}

	if !isCNPJFormationValidWithDV(cnpj) {
		return ErrCNPJNotValid
	}

	givenDV := cnpj[cnpjSizeWithoutDV:]
	calculatedDV, _ := CalculateCNPJCheckDigits(cnpj[:cnpjSizeWithoutDV])

	if calculatedDV != givenDV {
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
	second := calculateDigit(baseCNPJ + fmt.Sprintf("%d", first))

	return fmt.Sprintf("%d%d", first, second), nil
}

func removeFormattingCharacters(cnpj string) string {
	replacer := strings.NewReplacer(".", "", "/", "", "-", "")
	return replacer.Replace(trim(cnpj))
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

func isCNPJFormationValidWithDV(cnpj string) bool {
	return len(cnpj) == cnpjSize &&
		hasValidBaseFormation(cnpj[:cnpjSizeWithoutDV]) &&
		hasOnlyDigits(cnpj[cnpjSizeWithoutDV:]) &&
		!hasOnlyZeroes(cnpj)
}

func hasValidBaseFormation(cnpj string) bool {
	return len(cnpj) == cnpjSizeWithoutDV && hasOnlyUppercaseLettersAndDigits(cnpj)
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

func hasOnlyDigits(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] < '0' || value[i] > '9' {
			return false
		}
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
