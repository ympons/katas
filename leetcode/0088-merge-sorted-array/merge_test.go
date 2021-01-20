package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, merge(test.nums1, test.m, test.nums2, test.n))
	}
}

var tests = []struct {
	nums1, nums2 []int
	m, n         int
	expected     []int
}{
	{
		nums1:    []int{1, 2, 3, 0, 0, 0},
		m:        3,
		nums2:    []int{2, 5, 6},
		n:        3,
		expected: []int{1, 2, 2, 3, 5, 6},
	},
	{
		nums1:    []int{0},
		m:        0,
		nums2:    []int{1},
		n:        1,
		expected: []int{1},
	},
	{
		nums1:    []int{1},
		m:        1,
		nums2:    []int{},
		n:        0,
		expected: []int{1},
	},
}
