package base64

import (
	stdbase64 "encoding/base64"
	"errors"
	"strings"
)

var (
	ErrBase64NotValid = errors.New("base64 is not valid")
)

var encodings = []*stdbase64.Encoding{
	stdbase64.StdEncoding.Strict(),
	stdbase64.RawStdEncoding.Strict(),
	stdbase64.URLEncoding.Strict(),
	stdbase64.RawURLEncoding.Strict(),
}

func IsBase64(value string) error {
	value = strings.TrimSpace(value)
	if value == "" || hasWhitespace(value) || !canDecodeBase64(value) {
		return ErrBase64NotValid
	}

	return nil
}

func hasWhitespace(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] == ' ' || value[i] == '\t' || value[i] == '\n' || value[i] == '\r' {
			return true
		}
	}

	return false
}

func canDecodeBase64(value string) bool {
	for _, encoding := range encodings {
		if _, err := encoding.DecodeString(value); err == nil {
			return true
		}
	}

	return false
}
