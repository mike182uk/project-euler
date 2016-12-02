package main

import "fmt"

func main() {
	fmt.Println(getSquareOfSumUntil(100) - getSumOfSquareUntil(100))
}

func getSumOfSquareUntil(x int) int {
	sum := 0

	for i := 1; i <= x; i++ {
		sum += i * i
	}

	return sum
}

func getSquareOfSumUntil(x int) int {
	sum := 0

	for i := 1; i <= x; i++ {
		sum += i
	}

	return sum * sum
}
