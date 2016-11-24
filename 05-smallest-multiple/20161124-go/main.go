package main

import "fmt"

func main() {
	fmt.Println(getLowestCommonMultiple(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20))
}

func getLowestCommonMultiple(x ...int) int {
	r := x[0]

	for i := 1; i < len(x); i++ {
		r = lcm(r, x[i])
	}

	return r
}

func lcm(x int, y int) int {
	return x * y / gcd(x, y)
}

func gcd(x int, y int) int {
	for y > 0 {
		z := y
		y = x % y
		x = z
	}

	return x
}
