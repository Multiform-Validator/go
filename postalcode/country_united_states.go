package postalcode

var unitedStatesRule = countryRule{
	validate: isUnitedStatesPostalCode,
}

func isUnitedStatesPostalCode(value string) bool {
	if len(value) == 5 {
		return hasOnlyDigits(value)
	}

	return len(value) == 10 &&
		hasOnlyDigits(value[:5]) &&
		value[5] == '-' &&
		hasOnlyDigits(value[6:])
}
