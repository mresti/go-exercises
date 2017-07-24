package section02

import (
	"strings"
)

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	tmp1 := strings.Replace(phone, " ", "",-1)
	tmp2 := strings.Replace(tmp1, "-", "",-1)
	result := "(" + tmp2[:3] + ") " + tmp2[3:6] + "-" + tmp2[6:]
	return result
}
