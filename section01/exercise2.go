package section01

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
	return true
}
