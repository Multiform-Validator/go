package telephone

var germanyRule = countryRule{
	code:       "49",
	minLength:  5,
	maxLength:  11,
	prefixFunc: isGermanyTelephonePrefixValid,
}

func isGermanyTelephonePrefixValid(value string) bool {
	return len(value) > 0 && value[0] >= '1' && value[0] <= '9'
}
