package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(getPrimeAtIndex(10001))
}

func getPrimeAtIndex(n int) int {
	cpi := 1
	i := 2

	for {
		i++

		bi := big.NewInt(int64(i))

		if bi.ProbablyPrime(10) {
			cpi++
		}

		if cpi == n {
			return i
		}
	}
}
