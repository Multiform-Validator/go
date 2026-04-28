package email

import (
	"errors"
	"strings"
)

var (
	ErrEmailNotValid = errors.New("email is not valid")
)

func IsEmail(email string) error {
	if !isEmailFormationValid(email) {
		return ErrEmailNotValid
	}

	return nil
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

	if !isLetter(local[0]) {
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
	case '.', '_', '%', '+', '-':
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

	seenLabels := make(map[string]struct{}, len(labels))
	for _, label := range labels {
		if !isDomainLabelValid(label) {
			return false
		}

		normalizedLabel := strings.ToLower(label)
		if _, ok := seenLabels[normalizedLabel]; ok {
			return false
		}
		seenLabels[normalizedLabel] = struct{}{}
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

	if !isLetter(label[0]) {
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
