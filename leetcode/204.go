package leetcode

import "math"

func countPrimes(n int) int {
	count := 0
	notPrime := make([]bool, n)
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if !notPrime[i] {
			for j := 2; i*j < n; j++ {
				notPrime[i*j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
		}
	}
	return count
}
