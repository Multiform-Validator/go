package telephone

var brazilRule = countryRule{
	code:       "55",
	minLength:  10,
	maxLength:  11,
	prefixFunc: isBrazilTelephonePrefixValid,
}

func isBrazilTelephonePrefixValid(value telephoneNumber, start int, length int) bool {
	if length < 2 {
		return false
	}

	first := value.digits[start]
	second := value.digits[start+1]
	return first >= '1' && first <= '9' && second >= '1' && second <= '9'
}
