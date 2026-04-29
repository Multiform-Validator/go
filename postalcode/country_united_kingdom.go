package postalcode

var unitedKingdomRule = countryRule{
	validate: isUnitedKingdomPostalCode,
}

func isUnitedKingdomPostalCode(value string) bool {
	length := len(value)
	if length < 5 || length > 8 {
		return false
	}

	index := 0
	if !isLetter(value[index]) {
		return false
	}
	index++

	if index < length && isLetter(value[index]) {
		index++
	}

	if index >= length || !isDigit(value[index]) {
		return false
	}
	index++

	if index < length && isAlphaNumeric(value[index]) {
		index++
	}

	if index < length && value[index] == ' ' {
		index++
	}

	return index+3 == length &&
		isDigit(value[index]) &&
		isLetter(value[index+1]) &&
		isLetter(value[index+2])
}
