package section02

import (
	"testing"
	"reflect"
)

func TestUniqueSimple(t *testing.T) {
	input := []int {1,25,3,5,4,3,25,1,4}
	want := []int {1,25,3,5,4}
	result := Unique(input)
	if !reflect.DeepEqual(result, want) {
		t.Error(`Unique({1,25,3,5,4,3,25,1,4}) = {1,25,3,5,4}`)
	}
}

