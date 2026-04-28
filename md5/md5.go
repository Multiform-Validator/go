package md5

import "errors"

const md5Size = 32

var (
	ErrMD5NotValid = errors.New("MD5 is not valid")
)

func IsMD5(value string) error {
	if !isMD5FormationValid(value) {
		return ErrMD5NotValid
	}

	return nil
}

func isMD5FormationValid(value string) bool {
	if len(value) != md5Size {
		return false
	}

	for i := 0; i < len(value); i++ {
		if !isHexDigit(value[i]) {
			return false
		}
	}

	return true
}

func isHexDigit(value byte) bool {
	return (value >= '0' && value <= '9') ||
		(value >= 'a' && value <= 'f') ||
		(value >= 'A' && value <= 'F')
}
