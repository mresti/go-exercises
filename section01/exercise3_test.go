package section01

import "testing"

func TestPalindromeSimple(t *testing.T) {
	if !IsPalindrome("racecar") {
		t.Error(`IsPalindrome("racecar") = false`)
	}
}

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
