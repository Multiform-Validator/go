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
		{
			name:      "uses long country alias fallback",
			value:     "+55 11 91234-5678",
			countries: []string{"Brazil"},
			wantErr:   nil,
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

func TestTelephoneHelperBranches(t *testing.T) {
	if hasTelephonePrefix(telephoneNumberForTest("1"), "55") {
		t.Fatal("hasTelephonePrefix() expected false for short value")
	}

	if hasTelephonePrefix(telephoneNumberForTest("65"), "55") {
		t.Fatal("hasTelephonePrefix() expected false for mismatched prefix")
	}

	if value, ok := normalizeShortCountry("ca"); !ok || value != "ca" {
		t.Fatalf("normalizeShortCountry() = %q, %v; want %q, %v", value, ok, "ca", true)
	}

	if _, ok := normalizeShortCountry("unknown"); ok {
		t.Fatal("normalizeShortCountry() expected false for unsupported short country")
	}
}

func TestCountryPrefixValidationBranches(t *testing.T) {
	tests := []struct {
		name string
		got  bool
		want bool
	}{
		{"brazil rejects values shorter than area code", validateTelephonePrefixForTest(isBrazilTelephonePrefixValid, "1"), false},
		{"brazil accepts valid area code", validateTelephonePrefixForTest(isBrazilTelephonePrefixValid, "11912345678"), true},
		{"china rejects empty value", validateTelephonePrefixForTest(isChinaTelephonePrefixValid, ""), false},
		{"china accepts non zero prefix", validateTelephonePrefixForTest(isChinaTelephonePrefixValid, "13123456789"), true},
		{"china rejects zero prefix", validateTelephonePrefixForTest(isChinaTelephonePrefixValid, "03123456789"), false},
		{"canada rejects invalid length", validateTelephonePrefixForTest(isCanadaTelephonePrefixValid, "212555019"), false},
		{"canada accepts valid nanp number", validateTelephonePrefixForTest(isCanadaTelephonePrefixValid, "4165550199"), true},
		{"canada rejects invalid area code", validateTelephonePrefixForTest(isCanadaTelephonePrefixValid, "1165550199"), false},
		{"france rejects invalid length", validateTelephonePrefixForTest(isFranceTelephonePrefixValid, "14268530"), false},
		{"france accepts non zero prefix", validateTelephonePrefixForTest(isFranceTelephonePrefixValid, "142685300"), true},
		{"france rejects zero prefix", validateTelephonePrefixForTest(isFranceTelephonePrefixValid, "042685300"), false},
		{"germany rejects empty value", validateTelephonePrefixForTest(isGermanyTelephonePrefixValid, ""), false},
		{"germany accepts non zero prefix", validateTelephonePrefixForTest(isGermanyTelephonePrefixValid, "30123456"), true},
		{"germany rejects zero prefix", validateTelephonePrefixForTest(isGermanyTelephonePrefixValid, "030123456"), false},
		{"india rejects invalid length", validateTelephonePrefixForTest(isIndiaTelephonePrefixValid, "987654321"), false},
		{"india accepts mobile prefix", validateTelephonePrefixForTest(isIndiaTelephonePrefixValid, "9876543210"), true},
		{"india rejects landline prefix", validateTelephonePrefixForTest(isIndiaTelephonePrefixValid, "5876543210"), false},
		{"italy rejects empty value", validateTelephonePrefixForTest(isItalyTelephonePrefixValid, ""), false},
		{"italy accepts landline prefix", validateTelephonePrefixForTest(isItalyTelephonePrefixValid, "066982"), true},
		{"italy accepts mobile prefix", validateTelephonePrefixForTest(isItalyTelephonePrefixValid, "3123456789"), true},
		{"italy rejects unsupported prefix", validateTelephonePrefixForTest(isItalyTelephonePrefixValid, "166982"), false},
		{"japan rejects empty value", validateTelephonePrefixForTest(isJapanTelephonePrefixValid, ""), false},
		{"japan accepts non zero prefix", validateTelephonePrefixForTest(isJapanTelephonePrefixValid, "9012345678"), true},
		{"japan rejects zero prefix", validateTelephonePrefixForTest(isJapanTelephonePrefixValid, "0123456789"), false},
		{"south korea rejects empty value", validateTelephonePrefixForTest(isSouthKoreaTelephonePrefixValid, ""), false},
		{"south korea accepts non zero prefix", validateTelephonePrefixForTest(isSouthKoreaTelephonePrefixValid, "1012345678"), true},
		{"south korea rejects zero prefix", validateTelephonePrefixForTest(isSouthKoreaTelephonePrefixValid, "012345678"), false},
		{"united kingdom rejects empty value", validateTelephonePrefixForTest(isUnitedKingdomTelephonePrefixValid, ""), false},
		{"united kingdom accepts geographic prefix", validateTelephonePrefixForTest(isUnitedKingdomTelephonePrefixValid, "2079460958"), true},
		{"united kingdom rejects unsupported prefix", validateTelephonePrefixForTest(isUnitedKingdomTelephonePrefixValid, "9079460958"), false},
		{"united states rejects invalid length", validateTelephonePrefixForTest(isUnitedStatesTelephonePrefixValid, "212555019"), false},
		{"united states accepts valid nanp number", validateTelephonePrefixForTest(isUnitedStatesTelephonePrefixValid, "2125550199"), true},
		{"united states rejects invalid area code", validateTelephonePrefixForTest(isUnitedStatesTelephonePrefixValid, "1125550199"), false},
		{"united states rejects invalid exchange code", validateTelephonePrefixForTest(isUnitedStatesTelephonePrefixValid, "2121550199"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("prefix validation = %v, want %v", tt.got, tt.want)
			}
		})
	}
}

func telephoneNumberForTest(value string) telephoneNumber {
	var number telephoneNumber
	number.length = len(value)
	copy(number.digits[:], value)
	return number
}

func validateTelephonePrefixForTest(fn func(telephoneNumber, int, int) bool, value string) bool {
	number := telephoneNumberForTest(value)
	return fn(number, 0, number.length)
}
