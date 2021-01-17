package leetcode

func tupleSameProduct(nums []int) int {
	m, n := map[int]int{}, len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			m[nums[i]*nums[j]]++
		}
	}
	answer := 0
	for _, c := range m {
		if c > 1 {
			answer += 4 * c * (c - 1)
		}
	}
	return answer
}
