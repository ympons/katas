package leetcode

func removeDuplicate(nums []int) int {
	i := 0
	for _, x := range nums {
		if i < 2 || x > nums[i-2] {
			nums[i] = x
			i++
		}
	}
	return i
}
