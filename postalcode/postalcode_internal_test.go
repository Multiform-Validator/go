package postalcode

import "testing"

func TestUnitedKingdomPostalCodeBranches(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{"rejects too short value", "A1AA", false},
		{"rejects non letter start", "1W1A 1AA", false},
		{"rejects missing required digit", "SWAA 1AA", false},
		{"accepts without optional outward suffix", "S1 1AA", true},
		{"rejects invalid inward digit", "SW1A AAA", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnitedKingdomPostalCode(tt.value); got != tt.want {
				t.Fatalf("isUnitedKingdomPostalCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeShortCountryAliases(t *testing.T) {
	tests := []struct {
		country string
		want    string
	}{
		{"BR", "br"},
		{"CA", "ca"},
		{"CH", "ch"},
		{"DE", "de"},
		{"ES", "es"},
		{"FR", "fr"},
		{"GB", "gb"},
		{"IT", "it"},
		{"JP", "jp"},
		{"NL", "nl"},
		{"UK", "gb"},
		{"US", "us"},
		{"ZA", "za"},
		{"BRA", "br"},
		{"CHE", "ch"},
		{"DEU", "de"},
		{"ESP", "es"},
		{"FRA", "fr"},
		{"GBR", "gb"},
		{"ITA", "it"},
		{"JPN", "jp"},
		{"NLD", "nl"},
		{"USA", "us"},
		{"ZAF", "za"},
	}

	for _, tt := range tests {
		t.Run(tt.country, func(t *testing.T) {
			got, ok := normalizeShortCountry(tt.country)
			if !ok || got != tt.want {
				t.Fatalf("normalizeShortCountry() = %q, %v; want %q, true", got, ok, tt.want)
			}
		})
	}

	if got, ok := normalizeShortCountry("ZZ"); ok || got != "" {
		t.Fatalf("normalizeShortCountry() = %q, %v; want empty, false", got, ok)
	}

	if got, ok := normalizeShortCountry("ZZZ"); ok || got != "" {
		t.Fatalf("normalizeShortCountry() = %q, %v; want empty, false", got, ok)
	}
}
