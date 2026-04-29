package telephone

import (
	"errors"
	"testing"
)

func TestIsTelephoneInternalBranches(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		countries []string
		wantErr   error
	}{
		{
			name:      "returns nil when country is blank",
			value:     "2125550199",
			countries: []string{"   "},
			wantErr:   nil,
		},
		{
			name:      "returns invalid when country-specific national digits are too long",
			value:     "21255501991",
			countries: []string{"US"},
			wantErr:   ErrTelephoneNotValid,
		},
		{
			name:    "returns invalid when only plus sign is present",
			value:   "+",
			wantErr: ErrTelephoneNotValid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsTelephone(tt.value, tt.countries...)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("IsTelephone() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestCountryPrefixValidationBranches(t *testing.T) {
	tests := []struct {
		name string
		got  bool
		want bool
	}{
		{"brazil rejects values shorter than area code", isBrazilTelephonePrefixValid([]byte("1")), false},
		{"brazil accepts valid area code", isBrazilTelephonePrefixValid([]byte("11912345678")), true},
		{"china rejects empty value", isChinaTelephonePrefixValid([]byte("")), false},
		{"china accepts non zero prefix", isChinaTelephonePrefixValid([]byte("13123456789")), true},
		{"china rejects zero prefix", isChinaTelephonePrefixValid([]byte("03123456789")), false},
		{"canada rejects invalid length", isCanadaTelephonePrefixValid([]byte("212555019")), false},
		{"canada accepts valid nanp number", isCanadaTelephonePrefixValid([]byte("4165550199")), true},
		{"canada rejects invalid area code", isCanadaTelephonePrefixValid([]byte("1165550199")), false},
		{"france rejects invalid length", isFranceTelephonePrefixValid([]byte("14268530")), false},
		{"france accepts non zero prefix", isFranceTelephonePrefixValid([]byte("142685300")), true},
		{"france rejects zero prefix", isFranceTelephonePrefixValid([]byte("042685300")), false},
		{"germany rejects empty value", isGermanyTelephonePrefixValid([]byte("")), false},
		{"germany accepts non zero prefix", isGermanyTelephonePrefixValid([]byte("30123456")), true},
		{"germany rejects zero prefix", isGermanyTelephonePrefixValid([]byte("030123456")), false},
		{"india rejects invalid length", isIndiaTelephonePrefixValid([]byte("987654321")), false},
		{"india accepts mobile prefix", isIndiaTelephonePrefixValid([]byte("9876543210")), true},
		{"india rejects landline prefix", isIndiaTelephonePrefixValid([]byte("5876543210")), false},
		{"italy rejects empty value", isItalyTelephonePrefixValid([]byte("")), false},
		{"italy accepts landline prefix", isItalyTelephonePrefixValid([]byte("066982")), true},
		{"italy accepts mobile prefix", isItalyTelephonePrefixValid([]byte("3123456789")), true},
		{"italy rejects unsupported prefix", isItalyTelephonePrefixValid([]byte("166982")), false},
		{"japan rejects empty value", isJapanTelephonePrefixValid([]byte("")), false},
		{"japan accepts non zero prefix", isJapanTelephonePrefixValid([]byte("9012345678")), true},
		{"japan rejects zero prefix", isJapanTelephonePrefixValid([]byte("0123456789")), false},
		{"south korea rejects empty value", isSouthKoreaTelephonePrefixValid([]byte("")), false},
		{"south korea accepts non zero prefix", isSouthKoreaTelephonePrefixValid([]byte("1012345678")), true},
		{"south korea rejects zero prefix", isSouthKoreaTelephonePrefixValid([]byte("012345678")), false},
		{"united kingdom rejects empty value", isUnitedKingdomTelephonePrefixValid([]byte("")), false},
		{"united kingdom accepts geographic prefix", isUnitedKingdomTelephonePrefixValid([]byte("2079460958")), true},
		{"united kingdom rejects unsupported prefix", isUnitedKingdomTelephonePrefixValid([]byte("9079460958")), false},
		{"united states rejects invalid length", isUnitedStatesTelephonePrefixValid([]byte("212555019")), false},
		{"united states accepts valid nanp number", isUnitedStatesTelephonePrefixValid([]byte("2125550199")), true},
		{"united states rejects invalid area code", isUnitedStatesTelephonePrefixValid([]byte("1125550199")), false},
		{"united states rejects invalid exchange code", isUnitedStatesTelephonePrefixValid([]byte("2121550199")), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("prefix validation = %v, want %v", tt.got, tt.want)
			}
		})
	}
}
