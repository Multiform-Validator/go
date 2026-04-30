package telephone

var franceRule = countryRule{
	code:       "33",
	minLength:  9,
	maxLength:  9,
	prefixFunc: isFranceTelephonePrefixValid,
}

func isFranceTelephonePrefixValid(value telephoneNumber, start int, length int) bool {
	return length == 9 && value.digits[start] >= '1' && value.digits[start] <= '9'
}
