package main

func main() {

}

func SimpleFunction(x, y int) int {
	if x == y { // == means testing whether two expressions are equal
		return 0
	} else if x > y {
		return 1
	} else { // we can infer x < y
		return -1
	}
}

func SameSign(x, y int) bool {
	if x*y >= 0 {
		return true
	}
	// if we make it here, we know that x *yb <= 0
	return false
}

// C is undefined variable when written this way, it is trapped in the else if part
//func PositiveDifference(a, b int) int {
//	if a == b {
//		return 0
//	} else if a > b {
//		var c int = a - b
//		c := a-b // no new variables on the left
//	} else { // a < b
//		c = b - a
//	}
//	return c
//}

//func PositiveDifference(a, b int) int {
//	var c int
//	if a == b {
//		c = 0
//	} else if a > b {
//		c = a - b
//	} else { // a < b
//		c = b - a
//	}
//	return c
//}

func PositiveDifference(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return a - b
	}
	return b - a
}

/*
 * Other conditions
 *
 * >, <, >=, <=, ==
 * "not equal to" !=
 * "not" !x (this evaluates to true when x is false and false when x is true)
 */
