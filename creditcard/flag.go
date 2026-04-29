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
	digits, count := extractCardPrefix(cardNumber)
	if count == 0 {
		return FlagUnknown
	}

	switch {
	case hasCardPrefix(digits, count, "4"):
		return FlagVisa
	case hasCardPrefixRange(digits, count, 2, 51, 55):
		return FlagMastercard
	case hasAnyCardPrefix(digits, count, "34", "37"):
		return FlagAmericanExpress
	case hasAnyCardPrefix(digits, count, "6011", "65"):
		return FlagDiscover
	case hasAnyCardPrefix(digits, count, "2131", "1800") || hasCardPrefixRange(digits, count, 5, 35000, 35999):
		return FlagJCB
	case hasCardPrefixRange(digits, count, 3, 300, 305) || hasAnyCardPrefix(digits, count, "36", "38"):
		return FlagDinersClub
	case hasCardPrefixRange(digits, count, 4, 5000, 5999) || hasAnyCardPrefix(digits, count, "6304", "6390") || hasCardPrefixRange(digits, count, 4, 6700, 6799):
		return FlagMaestro
	case hasAnyCardPrefix(digits, count, "62", "88"):
		return FlagUnionPay
	case hasCardPrefixRange(digits, count, 3, 637, 639):
		return FlagElo
	case hasAnyCardPrefix(digits, count, "3841", "60"):
		return FlagHipercard
	default:
		return FlagUnknown
	}
}

func extractCardPrefix(cardNumber string) ([5]byte, int) {
	var digits [5]byte
	count := 0
	for i := 0; i < len(cardNumber); i++ {
		c := cardNumber[i]
		if c < '0' || c > '9' {
			continue
		}
		if count < len(digits) {
			digits[count] = c
		}
		count++
	}

	if count > len(digits) {
		return digits, len(digits)
	}

	return digits, count
}

func hasAnyCardPrefix(cardNumber [5]byte, count int, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if hasCardPrefix(cardNumber, count, prefix) {
			return true
		}
	}

	return false
}

func hasCardPrefix(cardNumber [5]byte, count int, prefix string) bool {
	if count < len(prefix) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if cardNumber[i] != prefix[i] {
			return false
		}
	}

	return true
}

func hasCardPrefixRange(cardNumber [5]byte, count int, size int, min int, max int) bool {
	if count < size {
		return false
	}

	value := 0
	for i := 0; i < size; i++ {
		value = value*10 + int(cardNumber[i]-'0')
	}

	return value >= min && value <= max
}
