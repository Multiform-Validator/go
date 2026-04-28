package macaddress

import "errors"

var (
	ErrMACAddressNotValid = errors.New("MAC address is not valid")
)

func IsMACAddress(value string) error {
	cleaned := cleanMACAddress(value)
	if len(cleaned) != 12 {
		return ErrMACAddressNotValid
	}

	for i := 0; i < len(cleaned); i++ {
		if !isHexDigit(cleaned[i]) {
			return ErrMACAddressNotValid
		}
	}

	return nil
}

func cleanMACAddress(value string) string {
	cleaned := make([]byte, 0, len(value))
	for i := 0; i < len(value); i++ {
		c := value[i]
		if isAlphaNumeric(c) {
			cleaned = append(cleaned, c)
		}
	}

	return string(cleaned)
}

func isAlphaNumeric(value byte) bool {
	return (value >= '0' && value <= '9') ||
		(value >= 'A' && value <= 'Z') ||
		(value >= 'a' && value <= 'z')
}

func isHexDigit(value byte) bool {
	return (value >= '0' && value <= '9') ||
		(value >= 'A' && value <= 'F') ||
		(value >= 'a' && value <= 'f')
}
