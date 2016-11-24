package main

import "sort"

func GetLargestPalindromeFromProductOf(l int, h int) int {
	var p []int

	for i := l; i <= h; i++ {
		for j := i; j <= h; j++ {
			x := j * i

			if x == ReverseInt(x) {
				p = append(p, x)
			}
		}
	}

	sort.Sort(sort.IntSlice(p))

	return p[len(p)-1]
}

func ReverseInt(n int) int {
	r := 0
	rm := 0

	for n > 0 {
		rm = n % 10
		n = n / 10
		r = r*10 + rm
	}

	return r
}
