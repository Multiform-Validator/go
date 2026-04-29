package postalcode

import (
	"errors"
	"strings"
)

var (
	ErrPostalCodeNotValid            = errors.New("postal code is not valid")
	ErrPostalCodeCountryNotSupported = errors.New("postal code country is not supported")
)

type countryRule struct {
	validate func(string) bool
}

var countryAliases = map[string]string{
	"br":             "br",
	"bra":            "br",
	"brazil":         "br",
	"brasil":         "br",
	"ca":             "ca",
	"canada":         "ca",
	"ch":             "ch",
	"che":            "ch",
	"switzerland":    "ch",
	"suica":          "ch",
	"de":             "de",
	"deu":            "de",
	"germany":        "de",
	"alemanha":       "de",
	"es":             "es",
	"esp":            "es",
	"spain":          "es",
	"espanha":        "es",
	"fr":             "fr",
	"fra":            "fr",
	"france":         "fr",
	"gb":             "gb",
	"gbr":            "gb",
	"uk":             "gb",
	"united kingdom": "gb",
	"reino unido":    "gb",
	"it":             "it",
	"ita":            "it",
	"italy":          "it",
	"jp":             "jp",
	"jpn":            "jp",
	"japan":          "jp",
	"nl":             "nl",
	"nld":            "nl",
	"netherlands":    "nl",
	"holanda":        "nl",
	"us":             "us",
	"usa":            "us",
	"united states":  "us",
	"estados unidos": "us",
	"za":             "za",
	"zaf":            "za",
	"south africa":   "za",
	"africa do sul":  "za",
}

var countryRules = map[string]countryRule{
	"br": {validate: isBrazilPostalCode},
	"ca": {validate: isCanadaPostalCode},
	"ch": {validate: isFourDigitPostalCode},
	"de": {validate: isFiveDigitPostalCode},
	"es": {validate: isFiveDigitPostalCode},
	"fr": {validate: isFiveDigitPostalCode},
	"gb": {validate: isUnitedKingdomPostalCode},
	"it": {validate: isFiveDigitPostalCode},
	"jp": {validate: isJapanPostalCode},
	"nl": {validate: isNetherlandsPostalCode},
	"us": {validate: isUnitedStatesPostalCode},
	"za": {validate: isFourDigitPostalCode},
}

func IsPostalCode(value string, countries ...string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		return ErrPostalCodeNotValid
	}

	if len(countries) == 0 || strings.TrimSpace(countries[0]) == "" {
		if isPostalCodeSupportedByAnyCountry(value) {
			return nil
		}

		return ErrPostalCodeNotValid
	}

	country, ok := normalizeCountry(countries[0])
	if !ok {
		return ErrPostalCodeCountryNotSupported
	}

	if !countryRules[country].validate(value) {
		return ErrPostalCodeNotValid
	}

	return nil
}

func isPostalCodeSupportedByAnyCountry(value string) bool {
	for _, rule := range countryRules {
		if rule.validate(value) {
			return true
		}
	}

	return false
}

func isBrazilPostalCode(value string) bool {
	if len(value) == 8 {
		return hasOnlyDigits(value)
	}

	return len(value) == 9 &&
		hasOnlyDigits(value[:5]) &&
		value[5] == '-' &&
		hasOnlyDigits(value[6:])
}

func isCanadaPostalCode(value string) bool {
	if len(value) == 6 {
		return isLetter(value[0]) &&
			isDigit(value[1]) &&
			isLetter(value[2]) &&
			isDigit(value[3]) &&
			isLetter(value[4]) &&
			isDigit(value[5])
	}

	return len(value) == 7 &&
		isLetter(value[0]) &&
		isDigit(value[1]) &&
		isLetter(value[2]) &&
		(value[3] == ' ' || value[3] == '-') &&
		isDigit(value[4]) &&
		isLetter(value[5]) &&
		isDigit(value[6])
}

func isUnitedKingdomPostalCode(value string) bool {
	length := len(value)
	if length < 5 || length > 8 {
		return false
	}

	index := 0
	if !isLetter(value[index]) {
		return false
	}
	index++

	if index < length && isLetter(value[index]) {
		index++
	}

	if index >= length || !isDigit(value[index]) {
		return false
	}
	index++

	if index < length && isAlphaNumeric(value[index]) {
		index++
	}

	if index < length && value[index] == ' ' {
		index++
	}

	return index+3 == length &&
		isDigit(value[index]) &&
		isLetter(value[index+1]) &&
		isLetter(value[index+2])
}

func isJapanPostalCode(value string) bool {
	if len(value) == 7 {
		return hasOnlyDigits(value)
	}

	return len(value) == 8 &&
		hasOnlyDigits(value[:3]) &&
		value[3] == '-' &&
		hasOnlyDigits(value[4:])
}

func isNetherlandsPostalCode(value string) bool {
	if len(value) == 4 {
		return hasOnlyDigits(value)
	}

	if len(value) == 6 {
		return hasOnlyDigits(value[:4]) && isLetter(value[4]) && isLetter(value[5])
	}

	return len(value) == 7 &&
		hasOnlyDigits(value[:4]) &&
		value[4] == ' ' &&
		isLetter(value[5]) &&
		isLetter(value[6])
}

func isUnitedStatesPostalCode(value string) bool {
	if len(value) == 5 {
		return hasOnlyDigits(value)
	}

	return len(value) == 10 &&
		hasOnlyDigits(value[:5]) &&
		value[5] == '-' &&
		hasOnlyDigits(value[6:])
}

func isFourDigitPostalCode(value string) bool {
	return len(value) == 4 && hasOnlyDigits(value)
}

func isFiveDigitPostalCode(value string) bool {
	return len(value) == 5 && hasOnlyDigits(value)
}

func hasOnlyDigits(value string) bool {
	for i := 0; i < len(value); i++ {
		if !isDigit(value[i]) {
			return false
		}
	}

	return true
}

func isAlphaNumeric(value byte) bool {
	return isLetter(value) || isDigit(value)
}

func isLetter(value byte) bool {
	return (value >= 'A' && value <= 'Z') || (value >= 'a' && value <= 'z')
}

func isDigit(value byte) bool {
	return value >= '0' && value <= '9'
}

func normalizeCountry(country string) (string, bool) {
	country = strings.TrimSpace(country)
	if value, ok := normalizeShortCountry(country); ok {
		return value, true
	}

	country = strings.ToLower(strings.TrimSpace(country))
	country = strings.Join(strings.Fields(country), " ")
	value, ok := countryAliases[country]
	return value, ok
}

func normalizeShortCountry(country string) (string, bool) {
	switch {
	case equalFoldASCII(country, "br") || equalFoldASCII(country, "bra"):
		return "br", true
	case equalFoldASCII(country, "ca"):
		return "ca", true
	case equalFoldASCII(country, "ch") || equalFoldASCII(country, "che"):
		return "ch", true
	case equalFoldASCII(country, "de") || equalFoldASCII(country, "deu"):
		return "de", true
	case equalFoldASCII(country, "es") || equalFoldASCII(country, "esp"):
		return "es", true
	case equalFoldASCII(country, "fr") || equalFoldASCII(country, "fra"):
		return "fr", true
	case equalFoldASCII(country, "gb") || equalFoldASCII(country, "gbr") || equalFoldASCII(country, "uk"):
		return "gb", true
	case equalFoldASCII(country, "it") || equalFoldASCII(country, "ita"):
		return "it", true
	case equalFoldASCII(country, "jp") || equalFoldASCII(country, "jpn"):
		return "jp", true
	case equalFoldASCII(country, "nl") || equalFoldASCII(country, "nld"):
		return "nl", true
	case equalFoldASCII(country, "us") || equalFoldASCII(country, "usa"):
		return "us", true
	case equalFoldASCII(country, "za") || equalFoldASCII(country, "zaf"):
		return "za", true
	default:
		return "", false
	}
}

func equalFoldASCII(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		c := a[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		if c != b[i] {
			return false
		}
	}

	return true
}
