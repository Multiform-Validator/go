package email

import "testing"

func TestEmailCandidateScannerBranches(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		start     int
		wantMatch string
		wantNext  int
		wantOK    bool
	}{
		{"without at sign", "contact team", 0, "", len("contact team"), false},
		{"without local part", "@example.com", 0, "", 1, true},
		{"without domain part", "user@", 0, "", 5, true},
		{"with special local characters", "send to user!tag@example.com", 0, "user!tag@example.com", len("send to user!tag@example.com"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMatch, gotNext, gotOK := nextEmailCandidate(tt.value, tt.start)
			if gotMatch != tt.wantMatch || gotNext != tt.wantNext || gotOK != tt.wantOK {
				t.Fatalf("nextEmailCandidate() = (%q, %d, %v), want (%q, %d, %v)", gotMatch, gotNext, gotOK, tt.wantMatch, tt.wantNext, tt.wantOK)
			}
		})
	}
}

func TestEmailCandidateLocalCharacterBranches(t *testing.T) {
	tests := []struct {
		name  string
		value byte
		want  bool
	}{
		{"letter", 'a', true},
		{"digit", '1', true},
		{"special", '!', true},
		{"invalid", ',', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEmailCandidateLocalCharacter(tt.value); got != tt.want {
				t.Fatalf("isEmailCandidateLocalCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
