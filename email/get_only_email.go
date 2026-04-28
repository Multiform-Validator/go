package email

import (
	"regexp"
	"strings"
)

const NoEmailFound = "No email found"

var emailCandidatePattern = regexp.MustCompile("[A-Za-z0-9.!#$%&'*+/=?^_`{|}~-]+@[A-Za-z0-9.-]+")

type GetOnlyEmailOptions struct {
	CleanDomain  bool
	CleanDomains []string
	RepeatEmail  bool
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
