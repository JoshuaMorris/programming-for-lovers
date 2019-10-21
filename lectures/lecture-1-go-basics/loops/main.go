package main

import "fmt"

func main() {
	fmt.Println("Loops.")

	n := 5
	m := Factorial(n)

	fmt.Println(m) // 120

	// fmt.Println(Factorial(-100)) // should return undefined
	// once we add a check for negative numbers we see panic error

	fmt.Println(SumFirstNIntegers(10)) // 55

	//var i uint = 10
	//for ; i >= 0; i-- {
	//	fmt.Println(i) // repeats the underflow of a uint, infinite loop
	//}
}

// first, a factorial function

func Factorial(n int) int {
	// lets handle negative input
	if n < 0 {
		// Panic will print error and immediately end program
		panic("Error: Negative input given to factorial function.")
	}
	p := 1
	i := 1

	// while i <= n (Go doesn't have a "while" keyword, they use "for"
	// 		p = p*i
	//		i = i+1
	for i <= n {
		p *= i
		i++ // comment ths out will create an infinite loop
	}

	// I still lives here and has not been garbage collected yet

	return p
}

func AnotherFactorial(n int) int {
	if n < 0 {
		panic("Error: Negative input given to AnotherFactorial()!")
	}
	p := 1

	// for every integer i between 1 and n, p = P * i
	for i := 1; i <= n; i++ {
		p *= i
	}

	// for i := n; i >= 1; i-- { // not wrong
	//		p *= i
	//	}

	return p
}

// Exercise: Write a function in go using a while loop that takes an integer n
// and returns the sum of the first n positive integers

func SumFirstNIntegers(n int) int {
	sum := 0
	var i int = 1
	for i <= n {
		sum += i // sum = sum + 1
		i++      // i = i + 1, there is also i-- which means i = i - 1
	}

	return sum
}

func AnotherSumOfIntegers(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum = sum + i
	}

	return sum
}

// Exercise: Write a function SumEven that sums all even numbers up to and
// possibly including n

func SumEven(n int) int {
	sum := 0

	for i := 2; i <= n; i = i + 2 {
		sum += i
	}

	return sum
}
