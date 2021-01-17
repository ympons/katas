package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, findKthLargest(test.nums, test.k))
	}
}

func TestFindKthLargestHeap(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, findKthLargestHeap(test.nums, test.k))
	}
}

var tests = []struct {
	nums     []int
	k        int
	expected int
}{
	{
		nums:     []int{3, 2, 1, 5, 6, 4},
		k:        2,
		expected: 5,
	},
	{
		nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		k:        4,
		expected: 4,
	},
	{
		nums:     []int{10},
		k:        1,
		expected: 10,
	},
}
