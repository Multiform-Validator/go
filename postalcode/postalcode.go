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
	"br": brazilRule,
	"ca": canadaRule,
	"ch": switzerlandRule,
	"de": germanyRule,
	"es": spainRule,
	"fr": franceRule,
	"gb": unitedKingdomRule,
	"it": italyRule,
	"jp": japanRule,
	"nl": netherlandsRule,
	"us": unitedStatesRule,
	"za": southAfricaRule,
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
	switch len(country) {
	case 2:
		return normalizeTwoCharacterCountry(country)
	case 3:
		return normalizeThreeCharacterCountry(country)
	}

	return "", false
}

func normalizeTwoCharacterCountry(country string) (string, bool) {
	first := toLowerASCII(country[0])
	second := toLowerASCII(country[1])
	switch {
	case first == 'b' && second == 'r':
		return "br", true
	case first == 'c' && second == 'a':
		return "ca", true
	case first == 'c' && second == 'h':
		return "ch", true
	case first == 'd' && second == 'e':
		return "de", true
	case first == 'e' && second == 's':
		return "es", true
	case first == 'f' && second == 'r':
		return "fr", true
	case first == 'g' && second == 'b':
		return "gb", true
	case first == 'i' && second == 't':
		return "it", true
	case first == 'j' && second == 'p':
		return "jp", true
	case first == 'n' && second == 'l':
		return "nl", true
	case first == 'u' && second == 'k':
		return "gb", true
	case first == 'u' && second == 's':
		return "us", true
	case first == 'z' && second == 'a':
		return "za", true
	}

	return "", false
}

func normalizeThreeCharacterCountry(country string) (string, bool) {
	first := toLowerASCII(country[0])
	second := toLowerASCII(country[1])
	third := toLowerASCII(country[2])
	switch {
	case first == 'b' && second == 'r' && third == 'a':
		return "br", true
	case first == 'c' && second == 'h' && third == 'e':
		return "ch", true
	case first == 'd' && second == 'e' && third == 'u':
		return "de", true
	case first == 'e' && second == 's' && third == 'p':
		return "es", true
	case first == 'f' && second == 'r' && third == 'a':
		return "fr", true
	case first == 'g' && second == 'b' && third == 'r':
		return "gb", true
	case first == 'i' && second == 't' && third == 'a':
		return "it", true
	case first == 'j' && second == 'p' && third == 'n':
		return "jp", true
	case first == 'n' && second == 'l' && third == 'd':
		return "nl", true
	case first == 'u' && second == 's' && third == 'a':
		return "us", true
	case first == 'z' && second == 'a' && third == 'f':
		return "za", true
	}

	return "", false
}

func toLowerASCII(value byte) byte {
	if value >= 'A' && value <= 'Z' {
		return value + 'a' - 'A'
	}

	return value
}
