package main

import "fmt"

func main() {
	sum := 0

	for _, v := range GetFibonacciUntil(4000000) {
		if v%2 == 0 {
			sum += v
		}
	}

	fmt.Println(sum)
}
