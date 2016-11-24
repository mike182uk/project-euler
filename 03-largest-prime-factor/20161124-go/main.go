package main

import "fmt"

func main() {
	fmt.Println(getLargestPrimeFactorFor(600851475143))
}

func getLargestPrimeFactorFor(x int) int {
	lpf := 0

	for i := 2; i <= x; i++ {
		if x%i == 0 {
			lpf = i
			x = x / i
		}
	}

	return lpf
}
