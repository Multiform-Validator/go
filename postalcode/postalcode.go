package postalcode

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrPostalCodeNotValid            = errors.New("postal code is not valid")
	ErrPostalCodeCountryNotSupported = errors.New("postal code country is not supported")
)

type countryRule struct {
	pattern *regexp.Regexp
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
	"br": {pattern: regexp.MustCompile(`^\d{5}-?\d{3}$`)},
	"ca": {pattern: regexp.MustCompile(`^[A-Za-z]\d[A-Za-z][ -]?\d[A-Za-z]\d$`)},
	"ch": {pattern: regexp.MustCompile(`^\d{4}$`)},
	"de": {pattern: regexp.MustCompile(`^\d{5}$`)},
	"es": {pattern: regexp.MustCompile(`^\d{5}$`)},
	"fr": {pattern: regexp.MustCompile(`^\d{5}$`)},
	"gb": {pattern: regexp.MustCompile(`^[A-Za-z]{1,2}\d[A-Za-z\d]?\s?\d[A-Za-z]{2}$`)},
	"it": {pattern: regexp.MustCompile(`^\d{5}$`)},
	"jp": {pattern: regexp.MustCompile(`^\d{3}-?\d{4}$`)},
	"nl": {pattern: regexp.MustCompile(`^\d{4}\s?[A-Za-z]{2}$|^\d{4}$`)},
	"us": {pattern: regexp.MustCompile(`^\d{5}(-\d{4})?$`)},
	"za": {pattern: regexp.MustCompile(`^\d{4}$`)},
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

	if !countryRules[country].pattern.MatchString(value) {
		return ErrPostalCodeNotValid
	}

	return nil
}

func isPostalCodeSupportedByAnyCountry(value string) bool {
	for _, rule := range countryRules {
		if rule.pattern.MatchString(value) {
			return true
		}
	}

	return false
}

func normalizeCountry(country string) (string, bool) {
	country = strings.ToLower(strings.TrimSpace(country))
	country = strings.Join(strings.Fields(country), " ")
	value, ok := countryAliases[country]
	return value, ok
}
