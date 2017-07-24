package section02

import (
	"reflect"
	"testing"
)

func TestTopCharactersSimple(t *testing.T) {
	input := "Hello, 世界, Hallo Moin"

	want := map[rune]int{
		108: 4,
		111: 3,
		32:  3,
	}

	result := TopCharacters(input, 2)
	if !reflect.DeepEqual(result, want) {
		t.Error(`TopCharacters("Hello, 世界Hello, 世界",1) = map[108:4 111:3 32:3]`)
	}
}
