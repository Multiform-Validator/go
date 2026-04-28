package telephone

var unitedStatesRule = countryRule{
	code:       "1",
	minLength:  10,
	maxLength:  10,
	prefixFunc: isUnitedStatesTelephonePrefixValid,
}

func isUnitedStatesTelephonePrefixValid(value string) bool {
	if len(value) != 10 {
		return false
	}

	return value[0] >= '2' && value[0] <= '9' && value[3] >= '2' && value[3] <= '9'
}
