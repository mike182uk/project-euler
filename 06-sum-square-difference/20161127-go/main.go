package main

import "fmt"

func main() {
	fmt.Println(getSquareOfSumUntil(100) - getSumOfSquareUntil(100))
}

func getSumOfSquareUntil(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

func getSquareOfSumUntil(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}
