package telephone

import (
	"errors"
	"strings"
)

var (
	ErrTelephoneNotValid            = errors.New("telephone is not valid")
	ErrTelephoneCountryNotSupported = errors.New("telephone country is not supported")
)

type countryRule struct {
	code       string
	minLength  int
	maxLength  int
	prefixFunc func(telephoneNumber, int, int) bool
}

type telephoneNumber struct {
	digits [15]byte
	length int
}

var countryAliases = map[string]string{
	"1":              "us",
	"+1":             "us",
	"br":             "br",
	"bra":            "br",
	"brazil":         "br",
	"brasil":         "br",
	"55":             "br",
	"+55":            "br",
	"ca":             "ca",
	"canada":         "ca",
	"cn":             "cn",
	"chn":            "cn",
	"china":          "cn",
	"86":             "cn",
	"+86":            "cn",
	"de":             "de",
	"deu":            "de",
	"germany":        "de",
	"alemanha":       "de",
	"49":             "de",
	"+49":            "de",
	"fr":             "fr",
	"fra":            "fr",
	"france":         "fr",
	"33":             "fr",
	"+33":            "fr",
	"gb":             "gb",
	"gbr":            "gb",
	"uk":             "gb",
	"united kingdom": "gb",
	"reino unido":    "gb",
	"44":             "gb",
	"+44":            "gb",
	"in":             "in",
	"ind":            "in",
	"india":          "in",
	"91":             "in",
	"+91":            "in",
	"it":             "it",
	"ita":            "it",
	"italy":          "it",
	"39":             "it",
	"+39":            "it",
	"jp":             "jp",
	"jpn":            "jp",
	"japan":          "jp",
	"81":             "jp",
	"+81":            "jp",
	"kr":             "kr",
	"kor":            "kr",
	"south korea":    "kr",
	"korea":          "kr",
	"coreia do sul":  "kr",
	"82":             "kr",
	"+82":            "kr",
	"us":             "us",
	"usa":            "us",
	"united states":  "us",
	"estados unidos": "us",
}

var countryRules = map[string]countryRule{
	"br": brazilRule,
	"ca": canadaRule,
	"cn": chinaRule,
	"de": germanyRule,
	"fr": franceRule,
	"gb": unitedKingdomRule,
	"in": indiaRule,
	"it": italyRule,
	"jp": japanRule,
	"kr": southKoreaRule,
	"us": unitedStatesRule,
}

func IsTelephone(value string, countries ...string) error {
	number, hasInternationalPrefix, err := normalizeTelephone(value)
	if err != nil {
		return err
	}

	if number.length < 7 || number.length > 15 {
		return ErrTelephoneNotValid
	}

	if len(countries) == 0 || strings.TrimSpace(countries[0]) == "" {
		return nil
	}

	country, ok := normalizeCountry(countries[0])
	if !ok {
		return ErrTelephoneCountryNotSupported
	}

	rule := countryRules[country]
	nationalStart := 0
	nationalLength := number.length
	if hasTelephonePrefix(number, rule.code) && (hasInternationalPrefix || number.length > rule.maxLength) {
		nationalStart = len(rule.code)
		nationalLength -= nationalStart
	}

	if nationalLength < rule.minLength || nationalLength > rule.maxLength {
		return ErrTelephoneNotValid
	}

	if rule.prefixFunc != nil && !rule.prefixFunc(number, nationalStart, nationalLength) {
		return ErrTelephoneNotValid
	}

	return nil
}

func normalizeTelephone(value string) (telephoneNumber, bool, error) {
	var number telephoneNumber
	value = strings.TrimSpace(value)
	if value == "" {
		return number, false, ErrTelephoneNotValid
	}

	hasInternationalPrefix := false
	for i := 0; i < len(value); i++ {
		c := value[i]
		switch {
		case c >= '0' && c <= '9':
			if number.length == len(number.digits) {
				return number, hasInternationalPrefix, ErrTelephoneNotValid
			}
			number.digits[number.length] = c
			number.length++
		case c == '+':
			if i != 0 || hasInternationalPrefix {
				return number, false, ErrTelephoneNotValid
			}
			hasInternationalPrefix = true
		case c == ' ', c == '\t', c == '\n', c == '\r', c == '-', c == '.', c == '(', c == ')':
			continue
		default:
			return number, false, ErrTelephoneNotValid
		}
	}

	if number.length == 0 {
		return number, hasInternationalPrefix, ErrTelephoneNotValid
	}

	return number, hasInternationalPrefix, nil
}

func hasTelephonePrefix(value telephoneNumber, prefix string) bool {
	if value.length < len(prefix) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if value.digits[i] != prefix[i] {
			return false
		}
	}

	return true
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
	case 1:
		if country[0] == '1' {
			return "us", true
		}
	case 2:
		if country[0] == '+' && country[1] == '1' {
			return "us", true
		}
		return normalizeTwoCharacterCountry(country)
	case 3:
		if country[0] == '+' {
			return normalizeTelephoneCountryCode(country[1], country[2])
		}
		return normalizeThreeCharacterCountry(country)
	}

	return "", false
}

func normalizeTwoCharacterCountry(country string) (string, bool) {
	if value, ok := normalizeTelephoneCountryCode(country[0], country[1]); ok {
		return value, true
	}

	first := toLowerASCII(country[0])
	second := toLowerASCII(country[1])
	switch {
	case first == 'b' && second == 'r':
		return "br", true
	case first == 'c' && second == 'a':
		return "ca", true
	case first == 'c' && second == 'n':
		return "cn", true
	case first == 'd' && second == 'e':
		return "de", true
	case first == 'f' && second == 'r':
		return "fr", true
	case first == 'g' && second == 'b':
		return "gb", true
	case first == 'i' && second == 'n':
		return "in", true
	case first == 'i' && second == 't':
		return "it", true
	case first == 'j' && second == 'p':
		return "jp", true
	case first == 'k' && second == 'r':
		return "kr", true
	case first == 'u' && second == 'k':
		return "gb", true
	case first == 'u' && second == 's':
		return "us", true
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
	case first == 'c' && second == 'h' && third == 'n':
		return "cn", true
	case first == 'd' && second == 'e' && third == 'u':
		return "de", true
	case first == 'f' && second == 'r' && third == 'a':
		return "fr", true
	case first == 'g' && second == 'b' && third == 'r':
		return "gb", true
	case first == 'i' && second == 'n' && third == 'd':
		return "in", true
	case first == 'i' && second == 't' && third == 'a':
		return "it", true
	case first == 'j' && second == 'p' && third == 'n':
		return "jp", true
	case first == 'k' && second == 'o' && third == 'r':
		return "kr", true
	case first == 'u' && second == 's' && third == 'a':
		return "us", true
	}

	return "", false
}

func normalizeTelephoneCountryCode(first byte, second byte) (string, bool) {
	switch {
	case first == '3' && second == '3':
		return "fr", true
	case first == '3' && second == '9':
		return "it", true
	case first == '4' && second == '4':
		return "gb", true
	case first == '4' && second == '9':
		return "de", true
	case first == '5' && second == '5':
		return "br", true
	case first == '8' && second == '1':
		return "jp", true
	case first == '8' && second == '2':
		return "kr", true
	case first == '8' && second == '6':
		return "cn", true
	case first == '9' && second == '1':
		return "in", true
	}

	return "", false
}

func toLowerASCII(value byte) byte {
	if value >= 'A' && value <= 'Z' {
		return value + 'a' - 'A'
	}

	return value
}
