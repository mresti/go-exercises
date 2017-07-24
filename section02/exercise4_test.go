package section02

import (
	"testing"
)

func TestSliceProductSimple(t *testing.T) {
	input := []int {1,25,3,5,4}
	want := 1500
	result := SliceProduct(input)
	if result != want {
		t.Error(`SliceProduct({1,25,3,5,4}) != 1500`)
	}
}


