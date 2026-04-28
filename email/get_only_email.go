package email

import (
	"regexp"
	"strings"
)

const NoEmailFound = "No email found"

var emailCandidatePattern = regexp.MustCompile("[A-Za-z0-9.!#$%&'*+/=?^_`{|}~-]+@[A-Za-z0-9.-]+")

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
