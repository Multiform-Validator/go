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
		{"brazil rejects values shorter than area code", isBrazilTelephonePrefixValid("1"), false},
		{"brazil accepts valid area code", isBrazilTelephonePrefixValid("11912345678"), true},
		{"canada rejects invalid length", isCanadaTelephonePrefixValid("212555019"), false},
		{"canada accepts valid nanp number", isCanadaTelephonePrefixValid("4165550199"), true},
		{"canada rejects invalid area code", isCanadaTelephonePrefixValid("1165550199"), false},
		{"italy rejects empty value", isItalyTelephonePrefixValid(""), false},
		{"italy accepts landline prefix", isItalyTelephonePrefixValid("066982"), true},
		{"italy rejects unsupported prefix", isItalyTelephonePrefixValid("166982"), false},
		{"united kingdom rejects empty value", isUnitedKingdomTelephonePrefixValid(""), false},
		{"united kingdom accepts geographic prefix", isUnitedKingdomTelephonePrefixValid("2079460958"), true},
		{"united kingdom rejects unsupported prefix", isUnitedKingdomTelephonePrefixValid("9079460958"), false},
		{"united states rejects invalid length", isUnitedStatesTelephonePrefixValid("212555019"), false},
		{"united states accepts valid nanp number", isUnitedStatesTelephonePrefixValid("2125550199"), true},
		{"united states rejects invalid exchange code", isUnitedStatesTelephonePrefixValid("2121550199"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("prefix validation = %v, want %v", tt.got, tt.want)
			}
		})
	}
}
