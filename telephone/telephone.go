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
	prefixFunc func([]byte) bool
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

	digits := number.digits[:number.length]
	if len(digits) < 7 || len(digits) > 15 {
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
	nationalDigits := digits
	if hasTelephonePrefix(digits, rule.code) && (hasInternationalPrefix || len(digits) > rule.maxLength) {
		nationalDigits = digits[len(rule.code):]
	}

	if len(nationalDigits) < rule.minLength || len(nationalDigits) > rule.maxLength {
		return ErrTelephoneNotValid
	}

	if rule.prefixFunc != nil && !rule.prefixFunc(nationalDigits) {
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

func hasTelephonePrefix(value []byte, prefix string) bool {
	if len(value) < len(prefix) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if value[i] != prefix[i] {
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
	switch {
	case equalFoldASCII(country, "br") || equalFoldASCII(country, "bra") || country == "55" || country == "+55":
		return "br", true
	case equalFoldASCII(country, "ca"):
		return "ca", true
	case equalFoldASCII(country, "cn") || equalFoldASCII(country, "chn") || country == "86" || country == "+86":
		return "cn", true
	case equalFoldASCII(country, "de") || equalFoldASCII(country, "deu") || country == "49" || country == "+49":
		return "de", true
	case equalFoldASCII(country, "fr") || equalFoldASCII(country, "fra") || country == "33" || country == "+33":
		return "fr", true
	case equalFoldASCII(country, "gb") || equalFoldASCII(country, "gbr") || equalFoldASCII(country, "uk") || country == "44" || country == "+44":
		return "gb", true
	case equalFoldASCII(country, "in") || equalFoldASCII(country, "ind") || country == "91" || country == "+91":
		return "in", true
	case equalFoldASCII(country, "it") || equalFoldASCII(country, "ita") || country == "39" || country == "+39":
		return "it", true
	case equalFoldASCII(country, "jp") || equalFoldASCII(country, "jpn") || country == "81" || country == "+81":
		return "jp", true
	case equalFoldASCII(country, "kr") || equalFoldASCII(country, "kor") || country == "82" || country == "+82":
		return "kr", true
	case equalFoldASCII(country, "us") || equalFoldASCII(country, "usa") || country == "1" || country == "+1":
		return "us", true
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
