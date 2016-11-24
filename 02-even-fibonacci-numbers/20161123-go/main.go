package main

import "fmt"

func main() {
	sum := 0

	for _, v := range getFibonacciUntil(4000000) {
		if v%2 == 0 {
			sum += v
		}
	}

	fmt.Println(sum)
}

func getFibonacciUntil(x int) []int {
	f := []int{1, 2}

	for {
		nxt := f[len(f)-1] + f[len(f)-2]

		if nxt >= x {
			break
		}

		f = append(f, nxt)
	}

	return f
}
