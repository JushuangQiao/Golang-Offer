package leetcode

func numJewelsInStones(J string, S string) int {
	hasJ := make(map[string]bool)
	for _, j := range J {
		hasJ[string(j)] = true
	}
	ret := 0
	for _, s := range S {
		if hasJ[string(s)] {
			ret += 1
		}
	}
	return ret
}
