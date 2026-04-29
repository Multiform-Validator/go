package postalcode

var canadaRule = countryRule{
	validate: isCanadaPostalCode,
}

func isCanadaPostalCode(value string) bool {
	if len(value) == 6 {
		return isLetter(value[0]) &&
			isDigit(value[1]) &&
			isLetter(value[2]) &&
			isDigit(value[3]) &&
			isLetter(value[4]) &&
			isDigit(value[5])
	}

	return len(value) == 7 &&
		isLetter(value[0]) &&
		isDigit(value[1]) &&
		isLetter(value[2]) &&
		(value[3] == ' ' || value[3] == '-') &&
		isDigit(value[4]) &&
		isLetter(value[5]) &&
		isDigit(value[6])
}
