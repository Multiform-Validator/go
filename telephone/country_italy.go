package telephone

var italyRule = countryRule{
	code:       "39",
	minLength:  6,
	maxLength:  11,
	prefixFunc: isItalyTelephonePrefixValid,
}

func isItalyTelephonePrefixValid(value string) bool {
	if len(value) == 0 {
		return false
	}

	return value[0] == '0' || value[0] == '3'
}
