package telephone

var southKoreaRule = countryRule{
	code:       "82",
	minLength:  8,
	maxLength:  10,
	prefixFunc: isSouthKoreaTelephonePrefixValid,
}

func isSouthKoreaTelephonePrefixValid(value string) bool {
	return len(value) > 0 && value[0] >= '1' && value[0] <= '9'
}
