package leetcode

func longestMountain(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	answer := 0
	for i, j := 0, 0; i < n; i++ {
		j = i
		if j+1 < n && nums[j] < nums[j+1] {
			for j+1 < n && nums[j] < nums[j+1] {
				j++
			}
			if j+1 < n && nums[j] > nums[j+1] {
				for j+1 < n && nums[j] > nums[j+1] {
					j++
				}
				if answer < j-i+1 {
					answer = j - i + 1
				}
			}
		}
		if j > i {
			i = j - 1
		}
	}
	return answer
}
