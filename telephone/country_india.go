package telephone

var indiaRule = countryRule{
	code:       "91",
	minLength:  10,
	maxLength:  10,
	prefixFunc: isIndiaTelephonePrefixValid,
}

func isIndiaTelephonePrefixValid(value telephoneNumber, start int, length int) bool {
	return length == 10 && value.digits[start] >= '6' && value.digits[start] <= '9'
}
