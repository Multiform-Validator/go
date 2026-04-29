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
