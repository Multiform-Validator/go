package validate

import (
	"errors"
	"fmt"
	"unicode"
)

var (
	ErrPasswordTooLong              = errors.New("password is too long")
	ErrPasswordTooShort             = errors.New("password is too short")
	ErrPasswordUppercaseRequired    = errors.New("password requires at least one uppercase letter")
	ErrPasswordSpecialCharRequired  = errors.New("password requires at least one special character")
	ErrPasswordNumberRequired       = errors.New("password requires at least one number")
	ErrPasswordLetterRequired       = errors.New("password requires at least one letter")
	ErrPasswordLengthOptionsInvalid = errors.New("password length options are invalid")
)

// PasswordOptions configures length and character requirements for Password.
//
// The zero value only requires a non-empty password: MinLength defaults to 1
// and MaxLength defaults to no limit.
//
//	err := validate.Password("a")
//
// Configure the fields for stricter signup rules:
//
//	err := validate.Password("MyP@ssw0rd", validate.PasswordOptions{
//		MinLength:          8,
//		MaxLength:          64,
//		RequireUppercase:   true,
//		RequireSpecialChar: true,
//		RequireNumber:      true,
//		RequireLetter:      true,
//	})
type PasswordOptions struct {
	// MinLength sets the minimum number of bytes required. Zero means 1.
	MinLength int

	// MaxLength sets the maximum number of bytes allowed. Zero means unlimited.
	MaxLength int

	// RequireUppercase requires at least one Unicode uppercase character.
	RequireUppercase bool

	// RequireSpecialChar requires at least one of these ASCII characters:
	// ! @ # $ % ^ & * ( ) , . ? " : { } | < >
	//
	// Underscore is not considered a special character by this validator.
	RequireSpecialChar bool

	// RequireNumber requires at least one Unicode digit.
	RequireNumber bool

	// RequireLetter requires at least one Unicode letter.
	RequireLetter bool
}

func Password(value string, options ...PasswordOptions) error {
	option := getPasswordOptions(options)
	minLength, maxLength, err := getPasswordLengths(option)
	if err != nil {
		return err
	}

	if maxLength > 0 && len(value) > maxLength {
		return fmt.Errorf("%w: password must be between %d and %d characters", ErrPasswordTooLong, minLength, maxLength)
	}

	if len(value) < minLength {
		if maxLength == 0 {
			return fmt.Errorf("%w: password must be greater than %d characters", ErrPasswordTooShort, minLength)
		}

		return fmt.Errorf("%w: password must be between %d and %d characters", ErrPasswordTooShort, minLength, maxLength)
	}

	if option.RequireUppercase && !hasUppercase(value) {
		return ErrPasswordUppercaseRequired
	}

	if option.RequireSpecialChar && !hasSpecialChar(value) {
		return ErrPasswordSpecialCharRequired
	}

	if option.RequireNumber && !hasNumber(value) {
		return ErrPasswordNumberRequired
	}

	if option.RequireLetter && !hasLetter(value) {
		return ErrPasswordLetterRequired
	}

	return nil
}

func getPasswordOptions(options []PasswordOptions) PasswordOptions {
	if len(options) == 0 {
		return PasswordOptions{}
	}

	return options[0]
}

func getPasswordLengths(option PasswordOptions) (int, int, error) {
	minLength := option.MinLength
	if minLength == 0 {
		minLength = 1
	}

	maxLength := option.MaxLength
	if minLength < 1 || maxLength < 0 {
		return 0, 0, ErrPasswordLengthOptionsInvalid
	}

	if maxLength > 0 && minLength > maxLength {
		return 0, 0, ErrPasswordLengthOptionsInvalid
	}

	return minLength, maxLength, nil
}

func hasUppercase(value string) bool {
	for _, char := range value {
		if unicode.IsUpper(char) {
			return true
		}
	}

	return false
}

func hasNumber(value string) bool {
	for _, char := range value {
		if unicode.IsDigit(char) {
			return true
		}
	}

	return false
}

func hasLetter(value string) bool {
	for _, char := range value {
		if unicode.IsLetter(char) {
			return true
		}
	}

	return false
}

func hasSpecialChar(value string) bool {
	for _, char := range value {
		switch char {
		case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', ',', '.', '?', '"', ':', '{', '}', '|', '<', '>':
			return true
		}
	}

	return false
}
