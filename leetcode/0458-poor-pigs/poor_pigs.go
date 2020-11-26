package leetcode

import "math"

// O(1)  (t+1)^x >= n
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	t, n := float64(minutesToTest/minutesToDie+1), float64(buckets)
	x := math.Trunc(math.Log(n) / math.Log(t))
	if math.Pow(t, x) < n {
		return int(x) + 1
	}
	return int(x)
}
