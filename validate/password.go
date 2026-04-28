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

type PasswordOptions struct {
	MinLength          int
	MaxLength          int
	RequireUppercase   bool
	RequireSpecialChar bool
	RequireNumber      bool
	RequireLetter      bool
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
