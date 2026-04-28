package validate_test

import (
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/validate"
)

func TestEmail(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		options []validate.EmailOptions
		wantErr error
	}{
		{"valid email", "user@example.com", nil, nil},
		{"valid email with max length", "user@example.com", []validate.EmailOptions{{MaxLength: 20}}, nil},
		{"valid email with country", "user@example.com.br", []validate.EmailOptions{{Country: "br"}}, nil},
		{"valid email with country including dot", "user@example.com.br", []validate.EmailOptions{{Country: ".br"}}, nil},
		{"valid email with default valid domains", "user@gmail.com", []validate.EmailOptions{{ValidDomains: true}}, nil},
		{"valid email with custom domain including at sign", "user@company.dev", []validate.EmailOptions{{ValidDomainsList: []string{"@company.dev"}}}, nil},
		{"valid email with custom domain without at sign", "user@company.dev", []validate.EmailOptions{{ValidDomainsList: []string{"company.dev"}}}, nil},
		{"valid email with custom domain and spaces", "user@company.dev", []validate.EmailOptions{{ValidDomainsList: []string{" company.dev "}}}, nil},
		{"valid email with uppercase custom domain option", "user@company.dev", []validate.EmailOptions{{ValidDomainsList: []string{" COMPANY.DEV "}}}, nil},
		{"valid email with uppercase value and default valid domains", "User@Gmail.com", []validate.EmailOptions{{ValidDomains: true}}, nil},
		{"valid email with padded country option", "user@example.com.br", []validate.EmailOptions{{Country: "  BR  "}}, nil},
		{"invalid email empty value", "", nil, validate.ErrEmailEmpty},
		{"invalid email blank value", "   ", nil, validate.ErrEmailEmpty},
		{"invalid email format", "user.example.com", nil, validate.ErrEmailNotValid},
		{"invalid email too long", "user@example.com", []validate.EmailOptions{{MaxLength: 5}}, validate.ErrEmailTooLong},
		{"invalid email country", "user@example.com", []validate.EmailOptions{{Country: "br"}}, validate.ErrEmailCountryNotValid},
		{"invalid default domain", "user@example.com", []validate.EmailOptions{{ValidDomains: true}}, validate.ErrEmailDomainNotAllowed},
		{"invalid custom domain", "user@example.com", []validate.EmailOptions{{ValidDomainsList: []string{"company.dev"}}}, validate.ErrEmailDomainNotAllowed},
		{"invalid custom domain with empty and non matching domains", "user@example.com", []validate.EmailOptions{{ValidDomainsList: []string{"", "company.dev"}}}, validate.ErrEmailDomainNotAllowed},
		{"invalid max length", "user@example.com", []validate.EmailOptions{{MaxLength: -1}}, validate.ErrEmailMaxLengthNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Email(tt.value, tt.options...)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Email() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
