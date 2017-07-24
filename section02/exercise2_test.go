package section02

import "testing"

func TestAnagramSimple(t *testing.T) {
	if !Anagram("abcde", "edcba") {
		t.Error(`Anagram("epap","pepa") = false`)
	}
}

func TestAnagram(t *testing.T) {
	var tests = []struct {
		input1 string
		input2 string
		want   bool
	}{
		{"abcde", "edcba", true},
		{"apple", "pleap", true},
		{"abcde", "edcb", false},
		{"a", "界", false},
		{"Hello世界", "世界elloH", true},
	}
	for _, test := range tests {
		if got := Anagram(test.input1, test.input2); got != test.want {
			t.Errorf("Anagram(%v, %v) = %v", test.input1, test.input2, got)
		}
	}
}
