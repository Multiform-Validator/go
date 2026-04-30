package text

import (
	"bytes"
	"errors"
	"strings"
)

var (
	ErrValueNotBlank = errors.New("value is not blank")
	ErrValueNotEmpty = errors.New("value is not empty")
)

func IsBlank(value string) error {
	if strings.TrimSpace(value) != "" {
		return ErrValueNotBlank
	}

	return nil
}

func IsBlankBytes(value []byte) error {
	if len(bytes.TrimSpace(value)) != 0 {
		return ErrValueNotBlank
	}

	return nil
}

func IsEmpty(value string) error {
	if value != "" {
		return ErrValueNotEmpty
	}

	return nil
}

func IsEmptyBytes(value []byte) error {
	if len(value) != 0 {
		return ErrValueNotEmpty
	}

	return nil
}
