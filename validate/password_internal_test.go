package validate

import "testing"

func TestPasswordCharacterChecks(t *testing.T) {
	tests := []struct {
		name string
		got  bool
		want bool
	}{
		{"has uppercase", hasUppercase("abcD"), true},
		{"has unicode uppercase", hasUppercase("abcÁ"), true},
		{"does not have uppercase", hasUppercase("abcd"), false},
		{"has number", hasNumber("abc1"), true},
		{"has unicode number", hasNumber("abc٩"), true},
		{"does not have number", hasNumber("abcd"), false},
		{"has letter", hasLetter("123a"), true},
		{"has unicode letter", hasLetter("123ç"), true},
		{"does not have letter", hasLetter("1234"), false},
		{"has special char", hasSpecialChar("abc!"), true},
		{"does not have special char", hasSpecialChar("abcd"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("check = %v, want %v", tt.got, tt.want)
			}
		})
	}
}
