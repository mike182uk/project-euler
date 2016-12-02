package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(getPrimeAtIndex(10001))
}

func getPrimeAtIndex(x int) int {
	cpi := 1
	i := 2

	for {
		i++

		y := big.NewInt(int64(i))

		if y.ProbablyPrime(10) {
			cpi++
		}

		if cpi == x {
			return i
		}
	}
}
