package leetcode

func maximumUniqueSubarray(nums []int) int {
	n, M := len(nums), map[int]struct{}{}
	M[nums[0]] = struct{}{}
	answer := nums[0]
	sum, i, j := nums[0], 0, 1
	for i < n-1 && j < n {
		if _, ok := M[nums[j]]; ok {
			sum -= nums[i]
			delete(M, nums[i])
			i++
		} else {
			sum += nums[j]
			if answer < sum {
				answer = sum
			}
			M[nums[j]] = struct{}{}
			j++
		}
	}
	return answer
}
