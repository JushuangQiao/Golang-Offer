package leetcode

import (
	"math"
)

func titleToNumber(s string) int {
	ret := 0
	for i := len(s) - 1; i >= 0; i-- {
		ret += int(math.Pow(26, float64(len(s)-1-i))) * (int(s[i]-'A') + 1)
	}
	return ret
}
