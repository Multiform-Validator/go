package email

import (
	"strings"
)

const NoEmailFound = "No email found"

var defaultCleanDomains = []string{".com.br", ".com.io", ".com.pt", ".com", ".net", ".org", ".io", ".pt", ".br"}

// GetOnlyEmailOptions changes how emails are extracted from free-form text.
//
// By default, candidates are kept as found and duplicates are removed:
//
//	emails := email.GetOnlyEmails("a@site.com a@site.com")
//	// []string{"a@site.com"}
//
// Use CleanDomain or CleanDomains when extra text may be glued to the domain:
//
//	emails := email.GetOnlyEmails(
//		"contact john@gmail.comEXTRA",
//		email.GetOnlyEmailOptions{CleanDomain: true},
//	)
//	// []string{"john@gmail.com"}
type GetOnlyEmailOptions struct {
	// CleanDomains trims candidates to the longest matching custom domain.
	// When this slice is not empty it takes precedence over CleanDomain.
	//
	//	email.GetOnlyEmail("user@company.devXYZ",
	//		email.GetOnlyEmailOptions{CleanDomains: []string{".dev"}})
	//	// "user@company.dev"
	CleanDomains []string

	// CleanDomain trims candidates to a known domain suffix such as .com,
	// .com.br, .net, .org, .io, .pt, or .br.
	CleanDomain bool

	// RepeatEmail keeps duplicate addresses in the result. When false,
	// duplicates are removed after optional domain cleaning.
	RepeatEmail bool
}

func GetOnlyEmail(value string, options ...GetOnlyEmailOptions) string {
	option := getOnlyEmailOption(options)
	cleanDomains := getCleanDomains(option)

	for start := 0; start < len(value); {
		match, next, ok := nextEmailCandidate(value, start)
		if !ok {
			break
		}
		start = next

		if len(cleanDomains) > 0 {
			match = cleanEmailDomain(match, cleanDomains)
		}

		if IsEmail(match) == nil {
			return match
		}
	}

	return NoEmailFound
}

func GetOnlyEmails(value string, options ...GetOnlyEmailOptions) []string {
	option := getOnlyEmailOption(options)
	cleanDomains := getCleanDomains(option)
	emails := make([]string, 0)
	seen := make(map[string]struct{})

	for start := 0; start < len(value); {
		match, next, ok := nextEmailCandidate(value, start)
		if !ok {
			break
		}
		start = next

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

func nextEmailCandidate(value string, start int) (string, int, bool) {
	at := strings.IndexByte(value[start:], '@')
	if at == -1 {
		return "", len(value), false
	}
	at += start

	left := at
	for left > 0 && isEmailCandidateLocalCharacter(value[left-1]) {
		left--
	}

	right := at + 1
	for right < len(value) && isEmailCandidateDomainCharacter(value[right]) {
		right++
	}

	if left == at || right == at+1 {
		return "", at + 1, true
	}

	return value[left:right], right, true
}

func isEmailCandidateLocalCharacter(value byte) bool {
	if (value >= 'A' && value <= 'Z') || (value >= 'a' && value <= 'z') || (value >= '0' && value <= '9') {
		return true
	}

	switch value {
	case '.', '!', '#', '$', '%', '&', '\'', '*', '+', '/', '=', '?', '^', '_', '`', '{', '|', '}', '~', '-':
		return true
	}

	return false
}

func isEmailCandidateDomainCharacter(value byte) bool {
	return (value >= 'A' && value <= 'Z') ||
		(value >= 'a' && value <= 'z') ||
		(value >= '0' && value <= '9') ||
		value == '.' ||
		value == '-'
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
		return defaultCleanDomains
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
