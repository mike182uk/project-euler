package main

func GetFibonacciUntil(a int) []int {
	f := []int{1, 2}

	for {
		n := f[len(f)-1] + f[len(f)-2]

		if n >= a {
			break
		}

		f = append(f, n)
	}

	return f
}
