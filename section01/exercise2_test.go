package section01

import "testing"

func TestPrimeSimple(t *testing.T) {
	if !IsPrime(2) {
		t.Error(`IsPrime(2) = false`)
	}
}

func TestIsPrime(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{2, true},
		{7, true},
		{11, true},
		{24, false},
	}
	for _, test := range tests {
		if got := IsPrime(test.input); got != test.want {
			t.Errorf("IsPrime(%v) = %v", test.input, got)
		}
	}
}
