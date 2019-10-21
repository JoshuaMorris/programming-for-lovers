package main

import "fmt"

// Go won't read this line

/*
Everything here won't be read either (multi-line comment)

today: we'll see four variables types
int  : integers
uint : nonnegative integers
bool : true of false boolean variable
float64: decimal number

next time:
byte : single symbol
string : contiguous collection of symbols (words)
*/

func main() {
	fmt.Println("Let's get started!")

	var j int = 14 // j is an integer variable
	var x float64 = -2.3
	var yo string = "Hi" // default value for string is ""
	var u uint = 14
	var symbol byte = 'H'     // prints 721 for a value of 'H'
	var statement bool = true // defaults to false

	fmt.Println(j)         // 14
	fmt.Println(x)         // -2.3
	fmt.Println(yo)        // Hi
	fmt.Println(u)         // 14
	fmt.Println(symbol)    // 72
	fmt.Println(statement) // true

	// shorthand declaration
	i := -6     // automatically type int
	hi := "Yo " // go automatically give this type string
	k := 34     // automatically an int
	//secondStatement := true

	// we can do arithmetic on numeric variables

	fmt.Println((i + j) * 2 * k) // 544
	fmt.Println(2*x - 3.16)      // -7.76
	fmt.Println(hi + yo)         // Yo Hi

	fmt.Println(j / i) // 2
	fmt.Println(k / j) // Go view as integer division and discards the remainder

	// if we wanted the actual value k/j. we use type conversions.
	fmt.Println(float64(k) / float64(j))

	// not all type conversions will work
	//var ok bool = bool(0) // false Invalid type conversion
	//fmt.Println(ok)

	var p int = -187 // overflow is common when dealing with integers
	var s uint = uint(p)

	fmt.Println(s) // 18446744073709551429

	m := 9223372036854775807
	fmt.Println(m + 1) // -9223372036854775808

	//fmt.Println(j * x) // invalid operation
	fmt.Println(float64(j) * x) // -32.199999999999996

	fmt.Println(SumTwoInts(j, j))

	/*
	 * Go Demands that input variables have correct types
	 * w, z := DoubleAndDuplicate(m)
	 * fmt.Println(w, z)
	 */
	w, z := DoubleAndDuplicate(float64(m))
	fmt.Println(w, z)

	n := 17
	fmt.Println(AddOne(n)) // 18
	fmt.Println(n)         // 17 Doesn't change the original value

	o := SumTwoInts(n, p)
	fmt.Println(o) // -170

	fmt.Println(Pi()) // 3.14
	PrintHi()         // Hi
}

// first line of a function is the function signature
// func is a keyword, what follows is the name of the function
// the parenthesis wraps variables that are passed in,
// what is after the parenthesis is what is returned

// SumTwoInts takes two integers and returns their sum
func SumTwoInts(a int, b int) int {
	return a + b
}

func DoubleAndDuplicate(x float64) (float64, float64) {
	return 2.0 * x, 2.0 * x
}

func Pi() float64 { // Doesn't take imputs
	return 3.14 // wrong
}

// functions can also not return anything
func PrintHi() {
	fmt.Println("Hi")
}

func AddOne(n int) int {
	n += 1
	return n // n is destroyed at the end of the function execution
}

// func main is another example of a function that doesn't take inputs or return anything
