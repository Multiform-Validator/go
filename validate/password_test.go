package validate

import (
	"errors"
	"testing"
)

func TestPassword(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		options []PasswordOptions
		wantErr error
	}{
		{"valid password with default options", "a", nil, nil},
		{"valid password with all requirements", "MyP@ssw0rd", []PasswordOptions{{MinLength: 8, MaxLength: 20, RequireUppercase: true, RequireSpecialChar: true, RequireNumber: true, RequireLetter: true}}, nil},
		{"valid password exactly at min and max", "abcd", []PasswordOptions{{MinLength: 4, MaxLength: 4}}, nil},
		{"valid password with unicode uppercase number and letter", "Á٩", []PasswordOptions{{RequireUppercase: true, RequireNumber: true, RequireLetter: true}}, nil},
		{"valid password with each supported special char", `!@#$%^&*(),.?":{}|<>`, []PasswordOptions{{RequireSpecialChar: true}}, nil},
		{"invalid empty password", "", nil, ErrPasswordTooShort},
		{"invalid password too short with max length", "abc", []PasswordOptions{{MinLength: 8, MaxLength: 20}}, ErrPasswordTooShort},
		{"invalid password too long", "abcdefghijk", []PasswordOptions{{MinLength: 3, MaxLength: 5}}, ErrPasswordTooLong},
		{"invalid missing uppercase", "password1!", []PasswordOptions{{RequireUppercase: true}}, ErrPasswordUppercaseRequired},
		{"invalid missing special char", "Password1", []PasswordOptions{{RequireSpecialChar: true}}, ErrPasswordSpecialCharRequired},
		{"invalid missing number", "Password!", []PasswordOptions{{RequireNumber: true}}, ErrPasswordNumberRequired},
		{"invalid missing letter", "123456!", []PasswordOptions{{RequireLetter: true}}, ErrPasswordLetterRequired},
		{"invalid underscore is not configured as special char", "Password1_", []PasswordOptions{{RequireSpecialChar: true}}, ErrPasswordSpecialCharRequired},
		{"invalid negative min length", "password", []PasswordOptions{{MinLength: -1}}, ErrPasswordLengthOptionsInvalid},
		{"invalid negative max length", "password", []PasswordOptions{{MaxLength: -1}}, ErrPasswordLengthOptionsInvalid},
		{"invalid min greater than max", "password", []PasswordOptions{{MinLength: 10, MaxLength: 5}}, ErrPasswordLengthOptionsInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Password(tt.value, tt.options...)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Password() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestPasswordCharacterChecks(t *testing.T) {
	tests := []struct {
		name string
		got  bool
		want bool
	}{
		{"has uppercase", hasUppercase("abcD"), true},
		{"has unicode uppercase", hasUppercase("abcÁ"), true},
		{"does not have uppercase", hasUppercase("abcd"), false},
		{"has number", hasNumber("abc1"), true},
		{"has unicode number", hasNumber("abc٩"), true},
		{"does not have number", hasNumber("abcd"), false},
		{"has letter", hasLetter("123a"), true},
		{"has unicode letter", hasLetter("123ç"), true},
		{"does not have letter", hasLetter("1234"), false},
		{"has special char", hasSpecialChar("abc!"), true},
		{"does not have special char", hasSpecialChar("abcd"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("check = %v, want %v", tt.got, tt.want)
			}
		})
	}
}
