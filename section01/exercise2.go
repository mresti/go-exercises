package main

import "fmt"

func main() {
	// Integers
	i := 5         // the type of i is int
	fmt.Println(i) //  "5"

	var apples int32 = 10
	fmt.Println(apples) //  "10"
	var oranges int16 = 20
	fmt.Println(oranges) //  "20"
	//var compote int = apples + oranges // compile error
	var compote int = int(apples) + int(oranges) // correct way
	fmt.Println(compote)                         //  "30"

	// Floats
	f := 7.5       // the type of f is float65
	fmt.Println(f) //  "7.5"
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) //  "0 -0 +Inf -Inf NaN"

	// Booleans
	x := 5
	fmt.Println(x >= 4 && x < 6) // "true"

	// Strings

	// Immutable sequence of bytes
	s := "Hello world!"
	fmt.Println(len(s)) // "12"

	// s[i:j] yields new strings (can also use s[i:], s[:i] or s[:])
	s1 := "Golang"
	fmt.Println(len(s1)) // "6"
	fmt.Println(s1[i:])  // "g"
	fmt.Println(s1[:i])  // "Golan"
	fmt.Println(s1[:])   // "Golang"
}
