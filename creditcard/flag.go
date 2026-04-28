package creditcard

const (
	FlagVisa            = "Visa"
	FlagMastercard      = "Mastercard"
	FlagAmericanExpress = "American Express"
	FlagDiscover        = "Discover"
	FlagJCB             = "JCB"
	FlagDinersClub      = "Diners Club"
	FlagMaestro         = "Maestro"
	FlagUnionPay        = "UnionPay"
	FlagElo             = "Elo"
	FlagHipercard       = "Hipercard"
	FlagUnknown         = "Unknown"
)

func IdentifyFlagCard(cardNumber string) string {
	digits := extractCardDigits(cardNumber)
	if digits == "" {
		return FlagUnknown
	}

	switch {
	case hasCardPrefix(digits, "4"):
		return FlagVisa
	case hasCardPrefixRange(digits, 2, 51, 55):
		return FlagMastercard
	case hasAnyCardPrefix(digits, "34", "37"):
		return FlagAmericanExpress
	case hasAnyCardPrefix(digits, "6011", "65"):
		return FlagDiscover
	case hasAnyCardPrefix(digits, "2131", "1800") || hasCardPrefixRange(digits, 5, 35000, 35999):
		return FlagJCB
	case hasCardPrefixRange(digits, 3, 300, 305) || hasAnyCardPrefix(digits, "36", "38"):
		return FlagDinersClub
	case hasCardPrefixRange(digits, 4, 5000, 5999) || hasAnyCardPrefix(digits, "6304", "6390") || hasCardPrefixRange(digits, 4, 6700, 6799):
		return FlagMaestro
	case hasAnyCardPrefix(digits, "62", "88"):
		return FlagUnionPay
	case hasCardPrefixRange(digits, 3, 637, 639):
		return FlagElo
	case hasAnyCardPrefix(digits, "3841", "60"):
		return FlagHipercard
	default:
		return FlagUnknown
	}
}

func extractCardDigits(cardNumber string) string {
	digits := make([]byte, 0, len(cardNumber))
	for i := 0; i < len(cardNumber); i++ {
		c := cardNumber[i]
		if c >= '0' && c <= '9' {
			digits = append(digits, c)
		}
	}

	return string(digits)
}

func hasAnyCardPrefix(cardNumber string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if hasCardPrefix(cardNumber, prefix) {
			return true
		}
	}

	return false
}

func hasCardPrefix(cardNumber string, prefix string) bool {
	return len(cardNumber) >= len(prefix) && cardNumber[:len(prefix)] == prefix
}

func hasCardPrefixRange(cardNumber string, size int, min int, max int) bool {
	if len(cardNumber) < size {
		return false
	}

	value := 0
	for i := 0; i < size; i++ {
		value = value*10 + int(cardNumber[i]-'0')
	}

	return value >= min && value <= max
}
