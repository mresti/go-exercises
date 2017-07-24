package section02

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	result := 1
	for _, v := range e {
		result *= v
	}
	return result
}
