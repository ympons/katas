package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSum(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maxSum(test.nums1, test.nums2)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums1, nums2 []int
	expected     int
}{
	{
		nums1:    []int{2, 4, 5, 8, 10},
		nums2:    []int{4, 6, 8, 9},
		expected: 30,
	},
	{
		nums1:    []int{1, 3, 5, 7, 9},
		nums2:    []int{3, 5, 100},
		expected: 109,
	},
	{
		nums1:    []int{1, 2, 3, 4, 5},
		nums2:    []int{6, 7, 8, 9, 10},
		expected: 40,
	},
	{
		nums1:    []int{1, 4, 5, 8, 9, 11, 19},
		nums2:    []int{2, 3, 4, 11, 12},
		expected: 61,
	},
}
