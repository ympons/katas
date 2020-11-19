package leetcode

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	D := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		D[i] = nums[i]
		for j := i + 1; j < n; j++ {
			D[j] = nums[i] - D[j]
			if nums[j]-D[j-1] > D[j] {
				D[j] = nums[j] - D[j-1]
			}
		}
	}
	return D[n-1] >= 0
}
