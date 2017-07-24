package section02

import (
	"strings"
)

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	n := make(map[rune]int)
	for _, r := range s {
		result := strings.Count(s, string(r))
		if result > k {
			n[r] = result
		}
	}
	return n
}
