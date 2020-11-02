package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKSum(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := kSum(test.nums, test.target, test.k)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums      []int
	target, k int
	expected  [][]int
}{
	{
		nums:     []int{1, 0, -1, 0, -2, 2},
		target:   0,
		k:        4,
		expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
	},
	{
		nums:     []int{},
		target:   0,
		k:        5,
		expected: [][]int{},
	},
	{
		nums:     []int{-1, 0, 1, 2, -1, -4},
		target:   0,
		k:        3,
		expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		nums:     []int{-1, -5, -5, -3, 2, 5, 0, 4},
		target:   -7,
		k:        4,
		expected: [][]int{{-5, -5, -1, 4}, {-5, -3, -1, 2}},
	},
	{
		nums:   []int{-5, -5, 0, 1, 0, 0, -1, -1, 2, 3, 4, -6, 6},
		target: 10,
		k:      8,
		expected: [][]int{
			{-6, 0, 0, 1, 2, 3, 4, 6},
			{-5, -1, 0, 1, 2, 3, 4, 6},
			{-5, 0, 0, 0, 2, 3, 4, 6},
			{-1, -1, 0, 0, 0, 2, 4, 6},
			{-1, -1, 0, 0, 1, 2, 3, 6},
		},
	},
}
