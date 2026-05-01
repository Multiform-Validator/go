package base64

import (
	"errors"
	"strings"
)

var (
	ErrBase64NotValid = errors.New("base64 is not valid")
)

func IsBase64(value string) error {
	value = strings.TrimSpace(value)
	if !hasValidBase64Format(value) {
		return ErrBase64NotValid
	}

	return nil
}

func hasValidBase64Format(value string) bool {
	if len(value) == 0 || len(value)%4 == 1 {
		return false
	}

	padding := 0
	for i := len(value) - 1; i >= 0 && value[i] == '='; i-- {
		padding++
	}

	if padding > 2 || (padding > 0 && len(value)%4 != 0) {
		return false
	}

	dataEnd := len(value) - padding
	for i := 0; i < dataEnd; i++ {
		if base64Value(value[i]) < 0 {
			return false
		}
	}

	if padding > 0 {
		strictPaddingMasks := [3]int{0, 0x03, 0x0F}
		return base64Value(value[dataEnd-1])&strictPaddingMasks[padding] == 0
	}

	return true
}

func base64Value(value byte) int {
	switch {
	case value >= 'A' && value <= 'Z':
		return int(value - 'A')
	case value >= 'a' && value <= 'z':
		return int(value-'a') + 26
	case value >= '0' && value <= '9':
		return int(value-'0') + 52
	}

	if index := strings.IndexByte("+/-_", value); index >= 0 {
		return 62 + index/2
	}

	return -1
}
