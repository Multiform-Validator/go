package postalcode

var netherlandsRule = countryRule{
	validate: isNetherlandsPostalCode,
}

func isNetherlandsPostalCode(value string) bool {
	if len(value) == 4 {
		return hasOnlyDigits(value)
	}

	if len(value) == 6 {
		return hasOnlyDigits(value[:4]) && isLetter(value[4]) && isLetter(value[5])
	}

	return len(value) == 7 &&
		hasOnlyDigits(value[:4]) &&
		value[4] == ' ' &&
		isLetter(value[5]) &&
		isLetter(value[6])
}
