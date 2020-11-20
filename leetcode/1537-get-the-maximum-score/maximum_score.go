package leetcode

const limit = int(1e9) + 7

func maxSum(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	i, j, s1, s2, sum := 0, 0, 0, 0, 0
	for i < n && j < m {
		if nums1[i] < nums2[j] {
			s1 += nums1[i]
			i++
		} else if nums1[i] > nums2[j] {
			s2 += nums2[j]
			j++
		} else {
			sum = (sum + max(s1, s2) + nums1[i]) % limit
			s1, s2 = 0, 0
			i++
			j++
		}
	}

	for i < n {
		s1 += nums1[i]
		i++
	}

	for j < m {
		s2 += nums2[j]
		j++
	}

	return (sum + max(s1, s2)) % limit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
