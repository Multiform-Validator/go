package email_test

import (
	"errors"
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
		{"valid email with plus", "user+tag@example.com", nil},
		{"valid email with subdomain", "first.last@mail.example.com", nil},
		{"valid email with surrounding spaces", " user@example.com ", nil},
		{"invalid email empty value", "", email.ErrEmailNotValid},
		{"invalid email without at sign", "user.example.com", email.ErrEmailNotValid},
		{"invalid email with more than one at sign", "user@@example.com", email.ErrEmailNotValid},
		{"invalid email without local part", "@example.com", email.ErrEmailNotValid},
		{"invalid email without domain", "user@", email.ErrEmailNotValid},
		{"invalid email with consecutive local dots", "first..last@example.com", email.ErrEmailNotValid},
		{"invalid email with local dot at start", ".user@example.com", email.ErrEmailNotValid},
		{"invalid email without domain dot", "user@example", email.ErrEmailNotValid},
		{"invalid email with domain ending dot", "user@example.com.", email.ErrEmailNotValid},
		{"invalid email with domain label starting hyphen", "user@-example.com", email.ErrEmailNotValid},
		{"invalid email with domain label ending hyphen", "user@example-.com", email.ErrEmailNotValid},
		{"invalid email with domain label too long", "user@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com", email.ErrEmailNotValid},
		{"invalid email with unsupported domain character", "user@exa_mple.com", email.ErrEmailNotValid},
		{"invalid email with short tld", "user@example.c", email.ErrEmailNotValid},
		{"invalid email with numeric tld", "user@example.c0m", email.ErrEmailNotValid},
		{"invalid email with spaces", "us er@example.com", email.ErrEmailNotValid},
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

func TestIsEmailBytes(t *testing.T) {
	tests := []struct {
		name    string
		email   []byte
		wantErr error
	}{
		{"valid email bytes", []byte("user@example.com"), nil},
		{"valid email bytes with subdomain", []byte("first.last@mail.example.com"), nil},
		{"invalid email bytes empty value", []byte(""), email.ErrEmailNotValid},
		{"invalid email bytes nil value", nil, email.ErrEmailNotValid},
		{"invalid email bytes without at sign", []byte("user.example.com"), email.ErrEmailNotValid},
		{"invalid email bytes with consecutive local dots", []byte("first..last@example.com"), email.ErrEmailNotValid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := email.IsEmailBytes(tt.email)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsEmailBytes() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
