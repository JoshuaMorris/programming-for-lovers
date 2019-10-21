package gcd

//func main() {
//	x := 378
//	y := 273
//	d := TrivialGCD(x, y)
//	fmt.Println(d) // 21
//
//	x := 378202873
//	y := 273147834
//
//	start := time.Now()
//	result := TrivialGCD(x, y)
//	fmt.Println(result)
//	elapsed := time.Since(start)
//	log.Printf("gcd took %s", elapsed)
//
//	start1 := time.Now()
//	gcd(x, y)
//	elapsed1 := time.Since(start1)
//	log.Printf("gcd took %s", elapsed1)
//
//	start2 := time.Now()
//	EuclidGCD(x, y)
//	elapsed2 := time.Since(start2)
//	log.Printf("EuclidGCD took %s", elapsed2)
//}

/*
 * TrivialGCD(a,b)
 * 	d <- 1
 *	m <- Min2(a,b)
 *	for every integer p from 1 up to m
 *		if p is a divisor of a and b
 *			d <- p
 *	return d
 */

func TrivialGCD(a, b int) int {
	d := 1
	m := Min2(a, b)
	for p := 1; p <= m; p++ {
		// remainder operation Remainder(n, k) is n%k (14%3 = 1)
		// if p is a divisor of a {
		//if a%p == 0 {
		//	// if p is a divisor of b
		//	if b%p == 0 {
		//		// if i'm here, i know that the answer to both is "YES"
		//		d = p
		//	}
		//}
		if a%p == 0 && b%p == 0 { // if statement is true only if both are true
			d = p
		}
	}
	return d
}

func Min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// There is an "OR" of two statements as well; operastor is ||
// e || f is true if one or both of e and f are true. it's only false if
// both e and f are false.

func AnotherSumEven(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		// is i even?
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}

func GCD(a, b int) int {
	if b != 0 {
		return GCD(b, a%b)
	}
	return a
}

/*
EuclidGCD(a, b)
	While a not equal to b
		if a >b
			a = a-b
		else  // b > a
			b = b-a
	return a // or b
*/
func EuclidGCD(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}
