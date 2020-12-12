package leetcode

func getSumAbsoluteDifferences(nums []int) []int {
	sum, sub := 0, 0
	for _, x := range nums {
		sum += x
	}
	n := len(nums)
	L := make([]int, n)
	for i := 0; i < n; i++ {
		sum -= nums[i]
		L[i] = sub + i*nums[i] - (n-i-1)*nums[i] + sum
		sub -= nums[i]
	}
	return L
}
