package validate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Multiform-Validator/go/email"
)

const defaultEmailMaxLength = 400

var (
	ErrEmailEmpty             = errors.New("email cannot be empty")
	ErrEmailNotValid          = errors.New("email is not valid")
	ErrEmailTooLong           = errors.New("email is too long")
	ErrEmailCountryNotValid   = errors.New("email is not valid in the country")
	ErrEmailDomainNotAllowed  = errors.New("email domain is not allowed")
	ErrEmailMaxLengthNotValid = errors.New("email max length is not valid")
)

var defaultValidEmailDomains = []string{
	"@gmail.com",
	"@outlook.com",
	"@yahoo.com",
	"@icloud.com",
	"@hotmail.com",
	"@mail.ru",
	"@yandex.ru",
	"@gmx.com",
	"@zoho.com",
	"@protonmail.com",
	"@protonmail.ch",
}

type EmailOptions struct {
	MaxLength        int
	Country          string
	ValidDomains     bool
	ValidDomainsList []string
}

func Email(value string, options ...EmailOptions) error {
	option := getEmailOptions(options)
	maxLength, err := getEmailMaxLength(option)
	if err != nil {
		return err
	}

	value = strings.TrimSpace(value)
	if value == "" {
		return ErrEmailEmpty
	}

	if !hasValidEmailDomain(value, option) {
		return ErrEmailDomainNotAllowed
	}

	if email.IsEmail(value) != nil {
		return ErrEmailNotValid
	}

	if len(value) > maxLength {
		return fmt.Errorf("%w: email cannot be greater than %d characters", ErrEmailTooLong, maxLength)
	}

	if !hasValidEmailCountry(value, option.Country) {
		return ErrEmailCountryNotValid
	}

	return nil
}

func getEmailOptions(options []EmailOptions) EmailOptions {
	if len(options) == 0 {
		return EmailOptions{}
	}

	return options[0]
}

func getEmailMaxLength(option EmailOptions) (int, error) {
	if option.MaxLength < 0 {
		return 0, ErrEmailMaxLengthNotValid
	}

	if option.MaxLength == 0 {
		return defaultEmailMaxLength, nil
	}

	return option.MaxLength, nil
}

func hasValidEmailDomain(value string, option EmailOptions) bool {
	domains := option.ValidDomainsList
	if len(domains) == 0 && option.ValidDomains {
		domains = defaultValidEmailDomains
	}

	if len(domains) == 0 {
		return true
	}

	value = strings.ToLower(value)
	for _, domain := range domains {
		domain = normalizeEmailDomain(domain)
		if domain != "" && strings.HasSuffix(value, domain) {
			return true
		}
	}

	return false
}

func normalizeEmailDomain(domain string) string {
	domain = strings.ToLower(strings.TrimSpace(domain))
	if domain == "" {
		return ""
	}

	if !strings.HasPrefix(domain, "@") {
		domain = "@" + domain
	}

	return domain
}

func hasValidEmailCountry(value string, country string) bool {
	country = strings.ToLower(strings.TrimSpace(country))
	if country == "" {
		return true
	}

	country = strings.TrimPrefix(country, ".")
	return strings.HasSuffix(strings.ToLower(value), "."+country)
}
