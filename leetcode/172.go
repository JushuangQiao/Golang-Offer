package leetcode

import "math"

func trailingZeroes(n int) int {
	ret := 0
	carry := 1
	for {
		cnt := n / int(math.Pow(5, float64(carry)))
		if cnt == 0 {
			return ret
		}
		ret += cnt
		carry += 1
	}
}
