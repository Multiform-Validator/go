package port

import (
	"errors"
	"strconv"
	"strings"
)

const (
	minPort = 1
	maxPort = 65535
)

var (
	ErrPortNotValid               = errors.New("port is not valid")
	ErrPortMustBeBetween1And65535 = errors.New("port must be between 1 and 65535")
)

func IsPort(port string) error {
	port = strings.TrimSpace(port)
	if !hasOnlyDigits(port) {
		return ErrPortNotValid
	}

	value, _ := strconv.Atoi(port)
	return validatePortRange(value)
}

func IsPortBytes(port []byte) error {
	return IsPort(string(port))
}

func IsPortNumber(port int) error {
	return validatePortRange(port)
}

func validatePortRange(port int) error {
	if port < minPort || port > maxPort {
		return ErrPortMustBeBetween1And65535
	}

	return nil
}

func hasOnlyDigits(value string) bool {
	if len(value) == 0 {
		return false
	}

	for i := 0; i < len(value); i++ {
		if value[i] < '0' || value[i] > '9' {
			return false
		}
	}

	return true
}
