package telephone

var unitedKingdomRule = countryRule{
	code:       "44",
	minLength:  10,
	maxLength:  10,
	prefixFunc: isUnitedKingdomTelephonePrefixValid,
}

func isUnitedKingdomTelephonePrefixValid(value string) bool {
	if len(value) == 0 {
		return false
	}

	switch value[0] {
	case '1', '2', '3', '7', '8':
		return true
	}

	return false
}
