package main

func GetLargestPrimeFactorFor(a int) int {
	lpf := 0

	for i := 2; i <= a; i++ {
		if a%i == 0 {
			lpf = i
			a = a / i
		}
	}

	return lpf
}
