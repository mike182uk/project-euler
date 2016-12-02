package main

import "fmt"

func main() {
	fmt.Println(getSmallestMultiple(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20))
}

func getSmallestMultiple(x ...int) int {
	sm := 1

	for i := 1; i < len(x); i++ {
		sm = lcm(sm, x[i])
	}

	return sm
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
