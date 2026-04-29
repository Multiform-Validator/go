package telephone

var japanRule = countryRule{
	code:       "81",
	minLength:  9,
	maxLength:  10,
	prefixFunc: isJapanTelephonePrefixValid,
}

func isJapanTelephonePrefixValid(value []byte) bool {
	return len(value) > 0 && value[0] >= '1' && value[0] <= '9'
}
