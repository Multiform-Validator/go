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

// EmailOptions configures the stricter email validator.
//
// The zero value validates the email format, trims surrounding whitespace, and
// limits the address to 400 characters.
//
//	err := validate.Email("user@example.com")
//
// Combine fields when the application has stricter rules:
//
//	err := validate.Email("user@company.dev", validate.EmailOptions{
//		MaxLength:        80,
//		ValidDomainsList: []string{"company.dev"},
//	})
type EmailOptions struct {
	// MaxLength sets the maximum accepted length after trimming spaces. Zero
	// means 400; negative values return ErrEmailMaxLengthNotValid.
	MaxLength int

	// Country requires the email to end with the given country suffix.
	//
	// The value is case-insensitive and may include the leading dot: "br" and
	// ".br" both accept "user@example.com.br".
	Country string

	// ValidDomainsList restricts the address to these domains instead of the
	// default allowlist. Domains are case-insensitive, surrounding spaces are
	// ignored, and the leading "@" is optional.
	ValidDomainsList []string

	// ValidDomains restricts the address to the package default domain allowlist
	// when true, including common providers such as gmail.com and outlook.com.
	ValidDomains bool
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

	for _, domain := range domains {
		if hasEmailDomainSuffix(value, domain) {
			return true
		}
	}

	return false
}

func hasEmailDomainSuffix(value string, domain string) bool {
	domain = strings.TrimSpace(domain)
	if domain == "" {
		return false
	}

	if domain[0] == '@' {
		return len(value) >= len(domain) &&
			strings.EqualFold(value[len(value)-len(domain):], domain)
	}

	return len(value) > len(domain) &&
		value[len(value)-len(domain)-1] == '@' &&
		strings.EqualFold(value[len(value)-len(domain):], domain)
}

func hasValidEmailCountry(value string, country string) bool {
	country = strings.ToLower(strings.TrimSpace(country))
	if country == "" {
		return true
	}

	country = strings.TrimPrefix(country, ".")
	return strings.HasSuffix(strings.ToLower(value), "."+country)
}
