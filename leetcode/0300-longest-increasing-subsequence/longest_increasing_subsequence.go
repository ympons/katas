package leetcode

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	D, answer := make([]int, n), 1
	D[0] = 1
	for i := 1; i < n; i++ {
		max := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && D[j] > max {
				max = D[j]
			}
		}
		D[i] = max + 1
		if D[i] > answer {
			answer = D[i]
		}
	}

	return answer
}
