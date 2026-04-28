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
	prefixFunc func(string) bool
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
	digits, hasInternationalPrefix, err := normalizeTelephone(value)
	if err != nil {
		return err
	}

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
	if strings.HasPrefix(digits, rule.code) && (hasInternationalPrefix || len(digits) > rule.maxLength) {
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

func normalizeTelephone(value string) (string, bool, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", false, ErrTelephoneNotValid
	}

	var digits strings.Builder
	hasInternationalPrefix := false
	for i := 0; i < len(value); i++ {
		c := value[i]
		switch {
		case c >= '0' && c <= '9':
			digits.WriteByte(c)
		case c == '+':
			if i != 0 || hasInternationalPrefix {
				return "", false, ErrTelephoneNotValid
			}
			hasInternationalPrefix = true
		case c == ' ', c == '\t', c == '\n', c == '\r', c == '-', c == '.', c == '(', c == ')':
			continue
		default:
			return "", false, ErrTelephoneNotValid
		}
	}

	if digits.Len() == 0 {
		return "", hasInternationalPrefix, ErrTelephoneNotValid
	}

	return digits.String(), hasInternationalPrefix, nil
}

func normalizeCountry(country string) (string, bool) {
	country = strings.ToLower(strings.TrimSpace(country))
	country = strings.Join(strings.Fields(country), " ")
	value, ok := countryAliases[country]
	return value, ok
}
