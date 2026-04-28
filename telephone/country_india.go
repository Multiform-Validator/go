package telephone

var indiaRule = countryRule{
	code:       "91",
	minLength:  10,
	maxLength:  10,
	prefixFunc: isIndiaTelephonePrefixValid,
}

func isIndiaTelephonePrefixValid(value string) bool {
	return len(value) == 10 && value[0] >= '6' && value[0] <= '9'
}
