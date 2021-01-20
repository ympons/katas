package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i, j, v := m-1, n-1, 0
	for k := m + n - 1; k >= 0; k-- {
		if i >= 0 || j >= 0 {
			if j < 0 || i >= 0 && nums1[i] > nums2[j] {
				v = nums1[i]
				i--
			} else {
				v = nums2[j]
				j--
			}
		}
		nums1[k] = v
	}
	return nums1
}
