package leetcode

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	for i, n := 0, 0; i < k; i++ {
		n = (n*10 + 1) % k
		if n == 0 {
			return i + 1
		}
	}
	return -1
}
