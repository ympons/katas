package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := fourSum(test.nums, test.target)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums     []int
	target   int
	expected [][]int
}{
	{
		nums:     []int{1, 0, -1, 0, -2, 2},
		target:   0,
		expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
	},
	{
		nums:     []int{},
		target:   0,
		expected: [][]int{},
	},
	{
		nums:     []int{-1, -5, -5, -3, 2, 5, 0, 4},
		target:   -7,
		expected: [][]int{{-5, -5, -1, 4}, {-5, -3, -1, 2}},
	},
}
