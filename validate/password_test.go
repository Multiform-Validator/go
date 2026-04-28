package validate_test

import (
	"errors"
	"testing"

	"github.com/Multiform-Validator/go/validate"
)

func TestPassword(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		options []validate.PasswordOptions
		wantErr error
	}{
		{"valid password with default options", "a", nil, nil},
		{"valid password with all requirements", "MyP@ssw0rd", []validate.PasswordOptions{{MinLength: 8, MaxLength: 20, RequireUppercase: true, RequireSpecialChar: true, RequireNumber: true, RequireLetter: true}}, nil},
		{"valid password exactly at min and max", "abcd", []validate.PasswordOptions{{MinLength: 4, MaxLength: 4}}, nil},
		{"valid password with unicode uppercase number and letter", "Á٩", []validate.PasswordOptions{{RequireUppercase: true, RequireNumber: true, RequireLetter: true}}, nil},
		{"valid password with each supported special char", `!@#$%^&*(),.?":{}|<>`, []validate.PasswordOptions{{RequireSpecialChar: true}}, nil},
		{"invalid empty password", "", nil, validate.ErrPasswordTooShort},
		{"invalid password too short with max length", "abc", []validate.PasswordOptions{{MinLength: 8, MaxLength: 20}}, validate.ErrPasswordTooShort},
		{"invalid password too long", "abcdefghijk", []validate.PasswordOptions{{MinLength: 3, MaxLength: 5}}, validate.ErrPasswordTooLong},
		{"invalid missing uppercase", "password1!", []validate.PasswordOptions{{RequireUppercase: true}}, validate.ErrPasswordUppercaseRequired},
		{"invalid missing special char", "Password1", []validate.PasswordOptions{{RequireSpecialChar: true}}, validate.ErrPasswordSpecialCharRequired},
		{"invalid missing number", "Password!", []validate.PasswordOptions{{RequireNumber: true}}, validate.ErrPasswordNumberRequired},
		{"invalid missing letter", "123456!", []validate.PasswordOptions{{RequireLetter: true}}, validate.ErrPasswordLetterRequired},
		{"invalid underscore is not configured as special char", "Password1_", []validate.PasswordOptions{{RequireSpecialChar: true}}, validate.ErrPasswordSpecialCharRequired},
		{"invalid negative min length", "password", []validate.PasswordOptions{{MinLength: -1}}, validate.ErrPasswordLengthOptionsInvalid},
		{"invalid negative max length", "password", []validate.PasswordOptions{{MaxLength: -1}}, validate.ErrPasswordLengthOptionsInvalid},
		{"invalid min greater than max", "password", []validate.PasswordOptions{{MinLength: 10, MaxLength: 5}}, validate.ErrPasswordLengthOptionsInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Password(tt.value, tt.options...)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Password() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
