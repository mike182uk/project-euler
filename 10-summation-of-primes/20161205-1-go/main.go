package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(getSumOfPrimesBelow(2000000))
}

func getSumOfPrimesBelow(x int) int {
	sum := 0

	for i := 0; i <= x; i++ {
		y := big.NewInt(int64(i))

		if y.ProbablyPrime(10) {
			sum += i
		}
	}

	return sum
}
