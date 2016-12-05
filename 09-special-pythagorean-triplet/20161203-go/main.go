package main

import "fmt"

func main() {
	var p int

	m := 2
	n := 1

	for {
		a, b, c := getPythagoreanTriplet(n, m)

		if a+b+c == 1000 {
			p = a * b * c
			break
		}

		n++

		if n == m {
			m++
			n = 1
		}
	}

	fmt.Println(p)
}

func getPythagoreanTriplet(n int, m int) (int, int, int) {
	x := m * m
	y := n * n

	a := x - y
	b := 2 * m * n
	c := x + y

	return a, b, c
}
