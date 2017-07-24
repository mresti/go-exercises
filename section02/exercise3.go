package section02

import (
	"sort"
)

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	var listEvens []int
	sort.Ints(e)
	for i := 0; i < len(e); i++ {
		if e[i]%2 != 0 {
			listEvens = append(listEvens, e[i])
		}
	}
	return listEvens
}
