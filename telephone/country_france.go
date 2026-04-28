package telephone

var franceRule = countryRule{
	code:       "33",
	minLength:  9,
	maxLength:  9,
	prefixFunc: isFranceTelephonePrefixValid,
}

func isFranceTelephonePrefixValid(value string) bool {
	return len(value) == 9 && value[0] >= '1' && value[0] <= '9'
}
