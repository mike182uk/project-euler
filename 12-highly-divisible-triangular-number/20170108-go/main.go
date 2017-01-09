package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getFirstTriangleNumberWithNfactors(500))
}

func getFirstTriangleNumberWithNfactors(n int) int {
	for i := 1; ; i++ {
		t := i * (i + 1) / 2

		if numOfDivisors(t) >= n {
			return t
		}
	}
}

func numOfDivisors(x int) int {
	d := 0

	for i := 1; float64(i) <= math.Sqrt(float64(x)); i++ {
		if x%i == 0 {
			d += 2
		}
	}

	return d
}
