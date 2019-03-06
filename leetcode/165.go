package leetcode

func compareVersion(version1 string, version2 string) int {
	k1, k2 := 0, 0
	for k1 < len(version1) && k2 < len(version2) {
		var v1 int
		for k1 < len(version1) {
			if version1[k1] == '.' {
				k1++
				break
			}
			v1 = 10*v1 + int(version1[k1]-'0')
			k1++
		}
		var v2 int
		for k2 < len(version2) {
			if version2[k2] == '.' {
				k2++
				break
			}
			v2 = 10*v2 + int(version2[k2]-'0')
			k2++
		}
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	for k1 < len(version1) {
		if !(version1[k1] == '.' || version1[k1] == '0') {
			return 1
		}
		k1++
	}
	for k2 < len(version2) {
		if !(version2[k2] == '.' || version2[k2] == '0') {
			return -1
		}
		k2++
	}
	return 0
}
