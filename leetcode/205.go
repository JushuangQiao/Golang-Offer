package leetcode

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	maps := make(map[string]string)
	mapt := make(map[string]string)
	for i := 0; i < len(s); i++ {
		a, b := string(s[i]), string(t[i])
		if maps[a] == "" {
			maps[a] = b
		}
		if mapt[b] == "" {
			mapt[b] = a
		}
		if maps[a] != b || mapt[b] != a {
			return false
		}
	}
	return true
}
