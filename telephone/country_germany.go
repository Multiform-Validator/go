package telephone

var germanyRule = countryRule{
	code:       "49",
	minLength:  5,
	maxLength:  11,
	prefixFunc: isGermanyTelephonePrefixValid,
}

func isGermanyTelephonePrefixValid(value telephoneNumber, start int, length int) bool {
	return length > 0 && value.digits[start] >= '1' && value.digits[start] <= '9'
}
