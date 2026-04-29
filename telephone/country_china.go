package telephone

var chinaRule = countryRule{
	code:       "86",
	minLength:  10,
	maxLength:  12,
	prefixFunc: isChinaTelephonePrefixValid,
}

func isChinaTelephonePrefixValid(value []byte) bool {
	return len(value) > 0 && value[0] >= '1' && value[0] <= '9'
}
