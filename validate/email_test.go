package validate

import (
	"errors"
	"testing"
)

func TestEmail(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		options []EmailOptions
		wantErr error
	}{
		{"valid email", "user@example.com", nil, nil},
		{"valid email with max length", "user@example.com", []EmailOptions{{MaxLength: 20}}, nil},
		{"valid email with country", "user@example.com.br", []EmailOptions{{Country: "br"}}, nil},
		{"valid email with country including dot", "user@example.com.br", []EmailOptions{{Country: ".br"}}, nil},
		{"valid email with default valid domains", "user@gmail.com", []EmailOptions{{ValidDomains: true}}, nil},
		{"valid email with custom domain including at sign", "user@company.dev", []EmailOptions{{ValidDomainsList: []string{"@company.dev"}}}, nil},
		{"valid email with custom domain without at sign", "user@company.dev", []EmailOptions{{ValidDomainsList: []string{"company.dev"}}}, nil},
		{"valid email with custom domain and spaces", "user@company.dev", []EmailOptions{{ValidDomainsList: []string{" company.dev "}}}, nil},
		{"valid email with uppercase custom domain option", "user@company.dev", []EmailOptions{{ValidDomainsList: []string{" COMPANY.DEV "}}}, nil},
		{"valid email with uppercase value and default valid domains", "User@Gmail.com", []EmailOptions{{ValidDomains: true}}, nil},
		{"valid email with padded country option", "user@example.com.br", []EmailOptions{{Country: "  BR  "}}, nil},
		{"invalid email empty value", "", nil, ErrEmailEmpty},
		{"invalid email blank value", "   ", nil, ErrEmailEmpty},
		{"invalid email format", "user.example.com", nil, ErrEmailNotValid},
		{"invalid email too long", "user@example.com", []EmailOptions{{MaxLength: 5}}, ErrEmailTooLong},
		{"invalid email country", "user@example.com", []EmailOptions{{Country: "br"}}, ErrEmailCountryNotValid},
		{"invalid default domain", "user@example.com", []EmailOptions{{ValidDomains: true}}, ErrEmailDomainNotAllowed},
		{"invalid custom domain", "user@example.com", []EmailOptions{{ValidDomainsList: []string{"company.dev"}}}, ErrEmailDomainNotAllowed},
		{"invalid custom domain with empty and non matching domains", "user@example.com", []EmailOptions{{ValidDomainsList: []string{"", "company.dev"}}}, ErrEmailDomainNotAllowed},
		{"invalid max length", "user@example.com", []EmailOptions{{MaxLength: -1}}, ErrEmailMaxLengthNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Email(tt.value, tt.options...)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Email() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmailDomainOptions(t *testing.T) {
	if !hasValidEmailDomain("user@example.com", EmailOptions{}) {
		t.Fatal("hasValidEmailDomain() expected true without domain options")
	}

	if hasValidEmailDomain("user@example.com", EmailOptions{ValidDomainsList: []string{""}}) {
		t.Fatal("hasValidEmailDomain() expected false with empty custom domain")
	}
}
