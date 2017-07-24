package section02

import (
	"sort"
)

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2){
		return false
	}

	runes1 := RuneSlice(s1)
	runes2 := RuneSlice(s2)

	// sort every list
	sort.Sort(RuneSlice(runes1))
	sort.Sort(RuneSlice(runes2))

	// check init values
	pos := 0
	matches := true

	for pos < len(runes1) && matches {
		if runes1[pos] == runes2[pos]{
			pos += 1
		} else{
			matches = false
		}
	}

	return matches
}
