package telephone

var brazilRule = countryRule{
	code:       "55",
	minLength:  10,
	maxLength:  11,
	prefixFunc: isBrazilTelephonePrefixValid,
}

func isBrazilTelephonePrefixValid(value []byte) bool {
	if len(value) < 2 {
		return false
	}

	first := value[0]
	second := value[1]
	return first >= '1' && first <= '9' && second >= '1' && second <= '9'
}
