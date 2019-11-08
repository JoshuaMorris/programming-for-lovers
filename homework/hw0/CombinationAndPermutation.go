/*
Homework 0 - Exercise 1
Exercise: Write and implement functions COMBINATION(n;k) and PERMUTATION(n;k)
computing the combination and permutation statistics.
*/

package main

import "fmt"

//Compute the factorial of n
func factorial(n int) int {
	f := 1
	for i := n; i > 0; i-- {
		f = f * i
	}
	return f
}

//Compute the permutation statistics
//n!/=((n-k)!)
func permutation(n int, k int) int {
	return factorial(n) / factorial(n-k)
}

// Compute the combination statistic
//C(n,k) = P(n,k)/k!.
func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}

func main() {
	fmt.Println(permutation(8, 4))
	fmt.Println(combination(8, 4))
}
