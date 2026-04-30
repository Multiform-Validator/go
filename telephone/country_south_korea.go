package telephone

var southKoreaRule = countryRule{
	code:       "82",
	minLength:  8,
	maxLength:  10,
	prefixFunc: isSouthKoreaTelephonePrefixValid,
}

func isSouthKoreaTelephonePrefixValid(value telephoneNumber, start int, length int) bool {
	return length > 0 && value.digits[start] >= '1' && value.digits[start] <= '9'
}
