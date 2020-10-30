package leetcode

const maxInt = 1<<31 - 1

func reverse(x int) int {
	r, c := 0, 0
	for x != 0 {
		c = x % 10
		x /= 10
		if limit := (maxInt - c) / 10; -limit > r || r > limit {
			return 0
		}
		r = r*10 + c
	}
	return r
}
