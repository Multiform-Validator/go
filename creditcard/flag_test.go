package creditcard_test

import (
	"testing"

	"github.com/Multiform-Validator/go/creditcard"
)

func TestIdentifyFlagCard(t *testing.T) {
	tests := []struct {
		name       string
		cardNumber string
		want       string
	}{
		{"visa", "4111 1111 1111 1111", creditcard.FlagVisa},
		{"mastercard", "5555555555554444", creditcard.FlagMastercard},
		{"american express", "378282246310005", creditcard.FlagAmericanExpress},
		{"discover with 6011", "6011111111111117", creditcard.FlagDiscover},
		{"discover with 65", "6500000000000000", creditcard.FlagDiscover},
		{"jcb with 2131", "213100000000000", creditcard.FlagJCB},
		{"jcb with 1800", "180000000000000", creditcard.FlagJCB},
		{"jcb with 35 range", "3528000000000000", creditcard.FlagJCB},
		{"diners club with 300 range", "30000000000000", creditcard.FlagDinersClub},
		{"diners club with 36", "36000000000000", creditcard.FlagDinersClub},
		{"diners club with 38", "38000000000000", creditcard.FlagDinersClub},
		{"maestro with 50 range", "5018000000000000", creditcard.FlagMaestro},
		{"maestro with 6304", "6304000000000000", creditcard.FlagMaestro},
		{"maestro with 6390", "6390000000000000", creditcard.FlagMaestro},
		{"maestro with 67 range", "6759000000000000", creditcard.FlagMaestro},
		{"unionpay with 62", "6200000000000000", creditcard.FlagUnionPay},
		{"unionpay with 88", "8800000000000000", creditcard.FlagUnionPay},
		{"elo", "6370000000000000", creditcard.FlagElo},
		{"3841 follows diners club precedence", "3841000000000000", creditcard.FlagDinersClub},
		{"hipercard with 60", "6062822686449791", creditcard.FlagHipercard},
		{"unknown", "7000000000000000", creditcard.FlagUnknown},
		{"unknown empty", "", creditcard.FlagUnknown},
		{"unknown without digits", "abcd", creditcard.FlagUnknown},
		{"short prefix does not panic", "3", creditcard.FlagUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := creditcard.IdentifyFlagCard(tt.cardNumber)
			if got != tt.want {
				t.Errorf("IdentifyFlagCard() = %q, want %q", got, tt.want)
			}
		})
	}
}
