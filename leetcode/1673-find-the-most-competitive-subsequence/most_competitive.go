package leetcode

func mostCompetitive(nums []int, k int) []int {
	n := len(nums)
	S := make([]int, 0, n)
	for i := 0; i < n; {
		if j := len(S); j > 0 && nums[i] < S[j-1] && i-j+k < n {
			S = S[0 : j-1]
		} else {
			S = append(S, nums[i])
			i++
		}
	}
	return S[0:k]
}
