package email

import (
	"errors"
	"regexp"
	"strings"
)

const NoEmailFound = "No email found"

var (
	ErrEmailNotValid = errors.New("email is not valid")
)

var emailCandidatePattern = regexp.MustCompile("[A-Za-z0-9.!#$%&'*+/=?^_`{|}~-]+@[A-Za-z0-9.-]+")

type GetOnlyEmailOptions struct {
	CleanDomain  bool
	CleanDomains []string
	RepeatEmail  bool
}

func IsEmail(email string) error {
	email = strings.TrimSpace(email)
	if !isEmailFormationValid(email) {
		return ErrEmailNotValid
	}

	return nil
}

func IsEmailBytes(email []byte) error {
	return IsEmail(string(email))
}

func GetOnlyEmail(value string, options ...GetOnlyEmailOptions) string {
	emails := GetOnlyEmails(value, options...)
	if len(emails) == 0 {
		return NoEmailFound
	}

	return emails[0]
}

func GetOnlyEmails(value string, options ...GetOnlyEmailOptions) []string {
	option := getOnlyEmailOption(options)
	cleanDomains := getCleanDomains(option)
	matches := emailCandidatePattern.FindAllString(value, -1)
	emails := make([]string, 0, len(matches))
	seen := make(map[string]struct{}, len(matches))

	for _, match := range matches {
		if len(cleanDomains) > 0 {
			match = cleanEmailDomain(match, cleanDomains)
		}

		if IsEmail(match) != nil {
			continue
		}

		if !option.RepeatEmail {
			if _, ok := seen[match]; ok {
				continue
			}
			seen[match] = struct{}{}
		}

		emails = append(emails, match)
	}

	return emails
}

func getOnlyEmailOption(options []GetOnlyEmailOptions) GetOnlyEmailOptions {
	if len(options) == 0 {
		return GetOnlyEmailOptions{}
	}

	return options[0]
}

func getCleanDomains(option GetOnlyEmailOptions) []string {
	if len(option.CleanDomains) > 0 {
		return option.CleanDomains
	}

	if option.CleanDomain {
		return []string{".com.br", ".com.io", ".com.pt", ".com", ".net", ".org", ".io", ".pt", ".br"}
	}

	return nil
}

func cleanEmailDomain(email string, domains []string) string {
	at := strings.LastIndexByte(email, '@')
	domain := email[at+1:]
	end := 0
	for _, allowedDomain := range domains {
		index := strings.Index(domain, allowedDomain)
		if index == -1 {
			continue
		}

		candidateEnd := index + len(allowedDomain)
		if candidateEnd > end {
			end = candidateEnd
		}
	}

	if end == 0 {
		return email
	}

	return email[:at+1+end]
}

func isEmailFormationValid(email string) bool {
	if len(email) == 0 || len(email) > 254 {
		return false
	}

	at := strings.IndexByte(email, '@')
	if at <= 0 || at != strings.LastIndexByte(email, '@') || at == len(email)-1 {
		return false
	}

	local := email[:at]
	domain := email[at+1:]

	return isLocalPartValid(local) && isDomainValid(domain)
}

func isLocalPartValid(local string) bool {
	if len(local) == 0 || len(local) > 64 || local[0] == '.' || local[len(local)-1] == '.' {
		return false
	}

	previousWasDot := false
	for i := 0; i < len(local); i++ {
		c := local[i]
		if c == '.' {
			if previousWasDot {
				return false
			}
			previousWasDot = true
			continue
		}

		if !isAllowedLocalCharacter(c) {
			return false
		}
		previousWasDot = false
	}

	return true
}

func isAllowedLocalCharacter(c byte) bool {
	if isLetter(c) || isDigit(c) {
		return true
	}

	switch c {
	case '!', '#', '$', '%', '&', '\'', '*', '+', '-', '/', '=', '?', '^', '_', '`', '{', '|', '}', '~':
		return true
	}

	return false
}

func isDomainValid(domain string) bool {
	if len(domain) == 0 || len(domain) > 253 || domain[0] == '.' || domain[len(domain)-1] == '.' {
		return false
	}

	labels := strings.Split(domain, ".")
	if len(labels) < 2 {
		return false
	}

	for _, label := range labels {
		if !isDomainLabelValid(label) {
			return false
		}
	}

	tld := labels[len(labels)-1]
	if len(tld) < 2 {
		return false
	}

	for i := 0; i < len(tld); i++ {
		if !isLetter(tld[i]) {
			return false
		}
	}

	return true
}

func isDomainLabelValid(label string) bool {
	if len(label) == 0 || len(label) > 63 {
		return false
	}

	if !isLetter(label[0]) && !isDigit(label[0]) {
		return false
	}

	if !isLetter(label[len(label)-1]) && !isDigit(label[len(label)-1]) {
		return false
	}

	for i := 0; i < len(label); i++ {
		c := label[i]
		if isLetter(c) || isDigit(c) || c == '-' {
			continue
		}

		return false
	}

	return true
}

func isLetter(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
