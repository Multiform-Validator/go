package telephone

var chinaRule = countryRule{
	code:       "86",
	minLength:  10,
	maxLength:  12,
	prefixFunc: isChinaTelephonePrefixValid,
}

func isChinaTelephonePrefixValid(value telephoneNumber, start int, length int) bool {
	return length > 0 && value.digits[start] >= '1' && value.digits[start] <= '9'
}
