package postalcode

var brazilRule = countryRule{
	validate: isBrazilPostalCode,
}

func isBrazilPostalCode(value string) bool {
	if len(value) == 8 {
		return hasOnlyDigits(value)
	}

	return len(value) == 9 &&
		hasOnlyDigits(value[:5]) &&
		value[5] == '-' &&
		hasOnlyDigits(value[6:])
}
