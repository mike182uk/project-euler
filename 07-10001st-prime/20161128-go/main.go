package main

import "fmt"

func main() {
	fmt.Println(getPrimeAtIndex(10001))
}

func getPrimeAtIndex(n int) int {
	c := 1
	p := 2

	for {
		i := p

		for {
			i++

			if isPrime(i) {
				p = i
				c++

				break
			}
		}

		if c == n {
			return i
		}
	}
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
