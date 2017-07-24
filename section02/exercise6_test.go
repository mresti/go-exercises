package section02

import (
	"testing"
	"reflect"
)

func TestInvertMapSimple(t *testing.T) {
	input := map[string]int{
		"apple":  1,
		"orange": 2,
	}

	want := map[int]string{
		1: "apple",
		2:"orange",
	}

	result := InvertMap(input)
	if !reflect.DeepEqual(result, want) {
		t.Error(`InvertMap(map[apple:1 orange:2]) = map[1:apple 2:orange]`)
	}
}

