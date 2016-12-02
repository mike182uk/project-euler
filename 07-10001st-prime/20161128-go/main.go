package main

import "fmt"

func main() {
	fmt.Println(getPrimeAtIndex(10001))
}

func getPrimeAtIndex(x int) int {
	cpi := 1
	p := 2

	for {
		i := p

		for {
			i++

			if isPrime(i) {
				cpi = i
				cpi++

				break
			}
		}

		if cpi == x {
			return i
		}
	}
}

func isPrime(x int) bool {
	if x == 2 {
		return true
	}

	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}
