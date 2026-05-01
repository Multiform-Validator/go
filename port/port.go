package port

import (
	"errors"
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
	value, ok := parsePort(port)
	if !ok {
		return ErrPortNotValid
	}

	return validatePortRange(value)
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

func parsePort(value string) (int, bool) {
	if len(value) == 0 {
		return 0, false
	}

	port := 0
	for i := 0; i < len(value); i++ {
		if value[i] < '0' || value[i] > '9' {
			return 0, false
		}

		port = port*10 + int(value[i]-'0')
		if port > maxPort {
			return port, true
		}
	}

	return port, true
}
