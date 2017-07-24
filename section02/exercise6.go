package section02

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	n := make(map[int]string)
	for k, v := range kv {
		n[v] = k
	}
	return n
}
