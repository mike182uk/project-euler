package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getLargestPalindromeFromProductOf(100, 999))
}

func getLargestPalindromeFromProductOf(x int, y int) int {
	var p []int

	for i := x; i <= y; i++ {
		for j := i; j <= y; j++ {
			z := j * i

			if z == reverseInt(z) {
				p = append(p, z)
			}
		}
	}

	sort.Sort(sort.IntSlice(p))

	return p[len(p)-1]
}

func reverseInt(x int) int {
	r := 0
	rm := 0

	for x > 0 {
		rm = x % 10
		x = x / 10
		r = r*10 + rm
	}

	return r
}
