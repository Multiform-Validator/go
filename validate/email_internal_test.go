package validate

import "testing"

func TestEmailDomainOptions(t *testing.T) {
	if !hasValidEmailDomain("user@example.com", EmailOptions{}) {
		t.Fatal("hasValidEmailDomain() expected true without domain options")
	}

	if hasValidEmailDomain("user@example.com", EmailOptions{ValidDomainsList: []string{""}}) {
		t.Fatal("hasValidEmailDomain() expected false with empty custom domain")
	}
}
