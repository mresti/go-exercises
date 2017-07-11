package main

import "fmt"

func main() {
	// Fizzbuzz function
	fmt.Println("Fizzbuzz function")
	fmt.Println(Fizzbuzz(3))
	fmt.Println(Fizzbuzz(5))
	fmt.Println(Fizzbuzz(15))
	fmt.Println(Fizzbuzz(11))

	// IsPrime function
	fmt.Println("IsPrime function")
	fmt.Println(IsPrime(2))  // "true"
	fmt.Println(IsPrime(11)) // "true"
	fmt.Println(IsPrime(24)) // "false"

	// Fizzbuzz function
	fmt.Println("IsPalindrome function")
	fmt.Println(IsPalindrome("Hello world!!"))
	fmt.Println(IsPalindrome("anna"))
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	result := ""
	if n%3 == 0 {
		result += "Fizz"
	}
	if n%5 == 0 {
		result += "Buzz"
	}
	return result
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	// check if n is a below of 2
	if n < 2 {
		return false
	}
	//check if n is a multiple of 2
	if n == 2 {
		return true
	}
	//check if n is a multiple of 2
	if n%2 == 0 {
		return false
	}
	//if not, then just check the odds
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return false
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
