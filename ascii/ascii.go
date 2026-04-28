package ascii

import (
	"errors"
)

var (
	ErrASCIINotValid = errors.New("ASCII is not valid")
)

func IsAscii(value string) error {
	return isAscii(value)

}

func IsAsciiBytes(value []byte) error {
	return isAscii(value)
}

func isAscii[T string | []byte](value T) error {
	for i := 0; i < len(value); i++ {
		if value[i] > 127 {
			return ErrASCIINotValid
		}
	}

	return nil
}
