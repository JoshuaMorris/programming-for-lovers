package primes

import "math"

// TrivialPrimeFinder takes an integer n and produces a boolean array of length
// n + 1 where primeArray[p] is true if p is prime and false otherwise
func TrivialPrimeFinder(n int) []bool {
	primeArray := make([]bool, n+1)
	for p := 2; p <= n; p++ {
		if IsPrime(p) == true {
			primeArray[p] = true
		}
	}
	return primeArray
}

func IsPrime(p int) bool {
	for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
		if p%k == 0 {
			return false
		}
	}
	return true
}

// SieveOfEratosthenes takes an integer n and returns a slice of n+1 booleans
// PrimeArray where primeArray[p] is true if p is prime and false otherwise
// it implements the Sieve of Eratosthenes approach
func SieveOfEratosthenes(n int) []bool {
	primeArray := make([]bool, n+1)

	// set everything to prime other than 0 and 1
	for k := 2; k <= n; k++ {
		primeArray[k] = true
	}

	// Now range over primeArray, and cross off multiples of first primes we see and iterate this process
	for p := 2; float64(p) <= math.Sqrt(float64(n)); p++ {
		if primeArray[p] == true {
			primeArray = CrossOffMultiples(primeArray, p)
		}
	}
	return primeArray
}

// CrossOfMultiples takes a boolean slice and an integer p. It crosses off
// multiples of p, meaning turning these multiples to false in the slice.
// It return the slice obtained as a result.
func CrossOffMultiples(primaryArray []bool, p int) []bool {
	n := len(primaryArray) - 1
	for k:= 2 *p; k <= n; k += p {
		primaryArray[k] = false
	}
	return primaryArray
}
