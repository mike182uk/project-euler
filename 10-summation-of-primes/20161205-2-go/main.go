package main

import "fmt"

func main() {
	fmt.Println(getSumOfPrimesBelow(2000000))
}

func getSumOfPrimesBelow(x int) int {
	sum := 0
	primes := generatePrimesUntil(x)

	for _, p := range primes {
		sum += p
	}

	return sum
}

func generatePrimesUntil(x int) []int {
	t := []bool{false, true}
	var p []int

	for i := 3; i <= x; i++ {
		t = append(t, true)
	}

	for i := 2; i < x; i++ {
		if t[i] == true {
			for j := i; j < x; j += i {
				t[j] = false
			}

			p = append(p, i)
		}
	}

	return p
}
