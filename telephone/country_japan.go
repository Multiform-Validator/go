package telephone

var japanRule = countryRule{
	code:       "81",
	minLength:  9,
	maxLength:  10,
	prefixFunc: isJapanTelephonePrefixValid,
}

func isJapanTelephonePrefixValid(value telephoneNumber, start int, length int) bool {
	return length > 0 && value.digits[start] >= '1' && value.digits[start] <= '9'
}
