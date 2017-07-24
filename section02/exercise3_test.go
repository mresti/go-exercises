package section02

import (
	"reflect"
	"testing"
)

func TestFindEvensSimple(t *testing.T) {
	input := []int{1, 25, 3, 5, 4}
	want := []int{1, 3, 5, 25}
	result := FindEvens(input)
	if !reflect.DeepEqual(result, want) {
		t.Error(`FindEvens({1,25,3,5,4}) = {1,3,5,25}`)
	}
}
