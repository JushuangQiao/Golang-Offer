package main

func nthUglyNumber(n int) int {
	var ugly []int
	ugly = append(ugly, 1)
	t2, t3, t5 := 0, 0, 0
	for n > len(ugly) {
		for ugly[t2]*2 <= ugly[len(ugly)-1] {
			t2 += 1
		}
		for ugly[t3]*3 <= ugly[len(ugly)-1] {
			t3 += 1
		}
		for ugly[t5]*5 <= ugly[len(ugly)-1] {
			t5 += 1
		}
		ugly = append(ugly, min(ugly[t2]*2, ugly[t3]*3, ugly[t5]*5))
	}
	return ugly[n-1]
}

func min(a, b, c int) int {
	if a <= b {
		b = a
	}
	if b <= c {
		return b
	}
	return c
}
