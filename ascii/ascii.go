package ascii

import "errors"

var (
	ErrASCIINotValid = errors.New("ASCII is not valid")
)

func IsAscii(value string) error {
	for i := 0; i < len(value); i++ {
		if value[i] > 127 {
			return ErrASCIINotValid
		}
	}

	return nil
}

func IsAsciiBytes(value []byte) error {
	for i := 0; i < len(value); i++ {
		if value[i] > 127 {
			return ErrASCIINotValid
		}
	}

	return nil
}
