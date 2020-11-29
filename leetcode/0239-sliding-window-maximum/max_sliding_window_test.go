package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindow(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maxSlidingWindow(test.nums, test.k)
		assert.Equal(test.expected, output)
	}
}

func TestMaxSlidingWindowPQ(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := maxSlidingWindowPQ(test.nums, test.k)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums     []int
	k        int
	expected []int
}{
	{
		nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
		k:        3,
		expected: []int{3, 3, 5, 5, 6, 7},
	},
	{
		nums:     []int{1},
		k:        1,
		expected: []int{1},
	},
	{
		nums:     []int{1, -1},
		k:        1,
		expected: []int{1, -1},
	},
	{
		nums:     []int{9, 11},
		k:        2,
		expected: []int{11},
	},
	{
		nums:     []int{4, -2},
		k:        2,
		expected: []int{4},
	},
}
