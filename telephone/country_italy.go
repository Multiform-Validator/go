package telephone

var italyRule = countryRule{
	code:       "39",
	minLength:  6,
	maxLength:  11,
	prefixFunc: isItalyTelephonePrefixValid,
}

func isItalyTelephonePrefixValid(value telephoneNumber, start int, length int) bool {
	if length == 0 {
		return false
	}

	return value.digits[start] == '0' || value.digits[start] == '3'
}
