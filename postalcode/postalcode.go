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
