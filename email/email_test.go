package email_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Multiform-Validator/go/email"
)

func TestIsEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr error
	}{
		{"valid email", "user@example.com", nil},
		{"valid email from TypeScript fixture", "foo@bar.com", nil},
		{"valid email with plus", "user+tag@example.com", nil},
		{"valid email with underscore", "user_name@example.com", nil},
		{"valid email with percent", "user%tag@example.com", nil},
		{"valid email with hyphen in domain label", "user@mail-server.example.com", nil},
		{"valid email with subdomain", "first.last@mail.example.com", nil},
		{"valid email with country code tld", "foo@bar.com.br", nil},
		{"valid email with uppercase letters", "Foo@Bar.com", nil},
		{"invalid email empty value", "", email.ErrEmailNotValid},
		{"invalid email without at sign", "user.example.com", email.ErrEmailNotValid},
		{"invalid email without at sign from TypeScript fixture", "foo.bar.com", email.ErrEmailNotValid},
		{"invalid email with more than one at sign", "user@@example.com", email.ErrEmailNotValid},
		{"invalid email with consecutive at signs", "foo@@bar.com", email.ErrEmailNotValid},
		{"invalid email with three at signs", "foo@@@bar.com", email.ErrEmailNotValid},
		{"invalid email with trailing at sign", "foo@bar.com@", email.ErrEmailNotValid},
		{"invalid email without local part", "@example.com", email.ErrEmailNotValid},
		{"invalid email without domain", "user@", email.ErrEmailNotValid},
		{"invalid email local starts with number", "1foo@bar.com", email.ErrEmailNotValid},
		{"invalid email domain starts with number", "foo@1bar.com", email.ErrEmailNotValid},
		{"invalid email with consecutive local dots", "first..last@example.com", email.ErrEmailNotValid},
		{"invalid email with consecutive local dots before at", "joao..@gmail.com", email.ErrEmailNotValid},
		{"invalid email with consecutive local dots from TypeScript fixture", "foo..bar@baz.com", email.ErrEmailNotValid},
		{"invalid email with local dot at start", ".user@example.com", email.ErrEmailNotValid},
		{"invalid email with local dot at end", "joao.@gmail.com", email.ErrEmailNotValid},
		{"invalid email without domain dot", "user@example", email.ErrEmailNotValid},
		{"invalid email with domain ending dot", "user@example.com.", email.ErrEmailNotValid},
		{"invalid email with domain ending consecutive dots", "foo@bar.com..", email.ErrEmailNotValid},
		{"invalid email with repeated domain label", "joaoaoao@gmail.com.com", email.ErrEmailNotValid},
		{"invalid email with repeated domain label case insensitive", "user@mail.Example.example", email.ErrEmailNotValid},
		{"invalid email with repeated country tld", "foo@bar.com.br.br", email.ErrEmailNotValid},
		{"invalid email with repeated domain and country tld", "foo@bar.com.com.br", email.ErrEmailNotValid},
		{"invalid email with repeated domain and repeated country tld", "foo@bar.com.com.br.br", email.ErrEmailNotValid},
		{"invalid email with consecutive domain dots", "joao@gmail..com", email.ErrEmailNotValid},
		{"invalid email with consecutive dots before domain", "foo...bar.com", email.ErrEmailNotValid},
		{"invalid email with consecutive dots before and after domain", "foo...@bar.com...", email.ErrEmailNotValid},
		{"invalid email with consecutive dots before and after domain and tld", "foo...@bar...com.", email.ErrEmailNotValid},
		{"invalid email with consecutive dots before and after at", "foo..@bar..com.", email.ErrEmailNotValid},
		{"invalid email with domain label starting hyphen", "user@-example.com", email.ErrEmailNotValid},
		{"invalid email with domain label ending hyphen", "user@example-.com", email.ErrEmailNotValid},
		{"invalid email with domain label too long", "user@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com", email.ErrEmailNotValid},
		{"invalid email with unsupported domain character", "user@exa_mple.com", email.ErrEmailNotValid},
		{"invalid email with domain label ending underscore", "user@example_.com", email.ErrEmailNotValid},
		{"invalid email with domain special character", "foo@bar!com", email.ErrEmailNotValid},
		{"invalid email with short tld", "user@example.c", email.ErrEmailNotValid},
		{"invalid email with short tld from TypeScript fixture", "foo@bar.c", email.ErrEmailNotValid},
		{"invalid email with numeric tld", "user@example.c0m", email.ErrEmailNotValid},
		{"invalid email with numeric tld label", "foo@bar.com.1", email.ErrEmailNotValid},
		{"invalid email with spaces", "us er@example.com", email.ErrEmailNotValid},
		{"invalid email with local spaces", "foo bar@baz.com", email.ErrEmailNotValid},
		{"invalid email with leading whitespace", " foo@bar.com", email.ErrEmailNotValid},
		{"invalid email with trailing whitespace", "foo@bar.com ", email.ErrEmailNotValid},
		{"invalid email local starts with comma", ",foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with exclamation", "!foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with hash", "#foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with dollar", "$foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with percent", "%foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with caret", "^foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with ampersand", "&foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with asterisk", "*foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with open parenthesis", "(foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with close parenthesis", ")foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with hyphen", "-foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with underscore", "_foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with equals", "=foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with plus", "+foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with open bracket", "[foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with close bracket", "]foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with open brace", "{foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with close brace", "}foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with semicolon", ";foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with colon", ":foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with single quote", "'foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with double quote", "\"foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with less than", "<foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with greater than", ">foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with question mark", "?foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with slash", "/foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with backslash", "\\foo@bar.com", email.ErrEmailNotValid},
		{"invalid email local starts with pipe", "|foo@bar.com", email.ErrEmailNotValid},
		{"invalid email with unsupported local special character", "foo!@bar.com", email.ErrEmailNotValid},
		{"invalid email with non ASCII local part", "usuário@example.com", email.ErrEmailNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := email.IsEmail(tt.email)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsEmail() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOnlyEmail(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		options []email.GetOnlyEmailOptions
		want    string
	}{
		{
			"returns first email",
			"Contact team: joao@empresa.com, maria@empresa.com, contato@empresa.com",
			[]email.GetOnlyEmailOptions{{}},
			"joao@empresa.com",
		},
		{
			"returns no email found when no email is present",
			"Contact team",
			[]email.GetOnlyEmailOptions{{}},
			email.NoEmailFound,
		},
		{
			"returns no email found without options",
			"Contact team",
			nil,
			email.NoEmailFound,
		},
		{
			"returns email without options",
			"Contact team:\talexa@google.com",
			nil,
			"alexa@google.com",
		},
		{
			"returns cleaned email with default domains",
			"Contact team:\talexa@google.com.br",
			[]email.GetOnlyEmailOptions{{CleanDomain: true}},
			"alexa@google.com.br",
		},
		{
			"returns cleaned email with custom domains",
			"Contact team:\talexa@google.custom",
			[]email.GetOnlyEmailOptions{{CleanDomains: []string{".custom"}}},
			"alexa@google.custom",
		},
		{
			"custom clean domains take precedence over default clean domain flag",
			"Contact team:\talexa@google.customEXTRA",
			[]email.GetOnlyEmailOptions{{CleanDomain: true, CleanDomains: []string{".custom"}}},
			"alexa@google.custom",
		},
		{
			"returns first email when repeated emails are allowed",
			"Contact team: john@gmail.com, john@gmail.com",
			[]email.GetOnlyEmailOptions{{RepeatEmail: true}},
			"john@gmail.com",
		},
		{
			"returns first cleaned email when repeated emails are allowed",
			"Contact team: john@gmail.comXTRA, alexa@gmail.comXTRA",
			[]email.GetOnlyEmailOptions{{CleanDomain: true, RepeatEmail: true}},
			"john@gmail.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := email.GetOnlyEmail(tt.value, tt.options...)
			if got != tt.want {
				t.Errorf("GetOnlyEmail() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestGetOnlyEmails(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		options []email.GetOnlyEmailOptions
		want    []string
	}{
		{
			"returns all emails",
			"Contact team: joao@empresa.com, maria@empresa.com, contato@empresa.com",
			[]email.GetOnlyEmailOptions{{}},
			[]string{"joao@empresa.com", "maria@empresa.com", "contato@empresa.com"},
		},
		{
			"returns cleaned emails",
			"Contact team: joao@empresa.com.br, maria@empresa.com.io, contato@empresa.com.pt jonyjony@gmail.comAwaodiawdoi",
			[]email.GetOnlyEmailOptions{{CleanDomain: true}},
			[]string{"joao@empresa.com.br", "maria@empresa.com.io", "contato@empresa.com.pt", "jonyjony@gmail.com"},
		},
		{
			"returns all emails without clean domain",
			"Contact team: john@gmail.com, jon2@gmail.com,",
			[]email.GetOnlyEmailOptions{{CleanDomain: false}},
			[]string{"john@gmail.com", "jon2@gmail.com"},
		},
		{
			"returns unique emails by default",
			"Contact team: joao@empresa.com, joao@empresa.com, joao@empresa.com",
			[]email.GetOnlyEmailOptions{{CleanDomain: false}},
			[]string{"joao@empresa.com"},
		},
		{
			"returns repeated emails when enabled",
			"Contact team: joao@empresa.com, joao@empresa.com, joao@empresa.com",
			[]email.GetOnlyEmailOptions{{CleanDomain: false, RepeatEmail: true}},
			[]string{"joao@empresa.com", "joao@empresa.com", "joao@empresa.com"},
		},
		{
			"returns all cleaned emails when repeated emails are allowed",
			"Contact team: john@gmail.comXTRA, alexa@gmail.comXTRA",
			[]email.GetOnlyEmailOptions{{CleanDomain: true, RepeatEmail: true}},
			[]string{"john@gmail.com", "alexa@gmail.com"},
		},
		{
			"deduplicates after cleaning domains",
			"Contact team: john@gmail.comXTRA, john@gmail.comOTHER",
			[]email.GetOnlyEmailOptions{{CleanDomains: []string{".com"}}},
			[]string{"john@gmail.com"},
		},
		{
			"skips invalid candidates",
			"Contact team: invalid@-example.com valid@example.com",
			nil,
			[]string{"valid@example.com"},
		},
		{
			"skips candidates with invalid cleaned domain",
			"Contact team: invalid@-example.comX valid@example.comX",
			[]email.GetOnlyEmailOptions{{CleanDomains: []string{".com"}}},
			[]string{"valid@example.com"},
		},
		{
			"keeps email when clean domain has no match",
			"Contact team: alexa@google.dev",
			[]email.GetOnlyEmailOptions{{CleanDomains: []string{".custom"}}},
			[]string{"alexa@google.dev"},
		},
		{
			"returns empty list when no email is present",
			"Contact team",
			nil,
			[]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := email.GetOnlyEmails(tt.value, tt.options...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOnlyEmails() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
