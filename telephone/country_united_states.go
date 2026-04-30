package telephone

var unitedStatesRule = countryRule{
	code:       "1",
	minLength:  10,
	maxLength:  10,
	prefixFunc: isUnitedStatesTelephonePrefixValid,
}

func isUnitedStatesTelephonePrefixValid(value telephoneNumber, start int, length int) bool {
	if length != 10 {
		return false
	}

	return value.digits[start] >= '2' && value.digits[start] <= '9' &&
		value.digits[start+3] >= '2' && value.digits[start+3] <= '9'
}
