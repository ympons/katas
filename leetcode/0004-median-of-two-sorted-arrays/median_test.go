package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := findMedianSortedArrays(test.nums1, test.nums2)
		assert.Equal(test.expected, output)
	}
}

func TestFindMedianSortedArraysOld(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := findMedianSortedArraysOld(test.nums1, test.nums2)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums1, nums2 []int
	expected     float64
}{
	{
		nums1:    []int{1, 3},
		nums2:    []int{2},
		expected: 2.00000,
	},
	{
		nums1:    []int{1, 2},
		nums2:    []int{3, 4},
		expected: 2.50000,
	},
	{
		nums1:    []int{0, 0},
		nums2:    []int{0, 0},
		expected: 0.00000,
	},
	{
		nums1:    []int{},
		nums2:    []int{1},
		expected: 1.00000,
	},
	{
		nums1:    []int{2},
		nums2:    []int{},
		expected: 2.00000,
	},
}
