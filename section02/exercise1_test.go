package section02

import "testing"

func TestParsePhoneSimple(t *testing.T) {
	result := ParsePhone("123-456-7890")
	if result != "(123) 456-7890" {
		t.Error(`ParsePhone(123456789) != (123) 456-7890`)
	}
}

func TestParsePhone(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"123-456-7890", "(123) 456-7890"},
		{"1 2 3 4 5 6 7 8 9 0", "(123) 456-7890"},
	}
	for _, test := range tests {
		if got := ParsePhone(test.input); got != test.want {
			t.Errorf("ParsePhone(%v) = %v", test.input, got)
		}
	}
}
