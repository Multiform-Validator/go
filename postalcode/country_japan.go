package postalcode

var japanRule = countryRule{
	validate: isJapanPostalCode,
}

func isJapanPostalCode(value string) bool {
	if len(value) == 7 {
		return hasOnlyDigits(value)
	}

	return len(value) == 8 &&
		hasOnlyDigits(value[:3]) &&
		value[3] == '-' &&
		hasOnlyDigits(value[4:])
}
