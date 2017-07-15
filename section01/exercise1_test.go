package section01

import "testing"

func TestFizzbuzzSimple(t *testing.T) {
	result := Fizzbuzz(3)
	if result != "Fizz" {
		t.Error(`Fizzbuzz(3) != Fizz`)
	}
}

func TestFizzbuzz(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{2, ""},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
	}
	for _, test := range tests {
		if got := Fizzbuzz(test.input); got != test.want {
			t.Errorf("Fizzbuzz(%v) = %v", test.input, got)
		}
	}
}
