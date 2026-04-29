package macaddress

import "errors"

var (
	ErrMACAddressNotValid = errors.New("MAC address is not valid")
)

func IsMACAddress(value string) error {
	digits := 0
	for i := 0; i < len(value); i++ {
		c := value[i]
		if !isAlphaNumeric(c) {
			continue
		}
		if !isHexDigit(c) {
			return ErrMACAddressNotValid
		}
		digits++
	}

	if digits != 12 {
		return ErrMACAddressNotValid
	}

	return nil
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
