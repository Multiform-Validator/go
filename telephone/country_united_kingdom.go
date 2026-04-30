package telephone

var unitedKingdomRule = countryRule{
	code:       "44",
	minLength:  10,
	maxLength:  10,
	prefixFunc: isUnitedKingdomTelephonePrefixValid,
}

func isUnitedKingdomTelephonePrefixValid(value telephoneNumber, start int, length int) bool {
	if length == 0 {
		return false
	}

	switch value.digits[start] {
	case '1', '2', '3', '7', '8':
		return true
	}

	return false
}
