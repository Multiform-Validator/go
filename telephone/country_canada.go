package telephone

var canadaRule = countryRule{
	code:       "1",
	minLength:  10,
	maxLength:  10,
	prefixFunc: isCanadaTelephonePrefixValid,
}

func isCanadaTelephonePrefixValid(value string) bool {
	if len(value) != 10 {
		return false
	}

	return value[0] >= '2' && value[0] <= '9' && value[3] >= '2' && value[3] <= '9'
}
