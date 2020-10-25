package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := twoSum(test.nums, test.target)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums     []int
	target   int
	expected []int
}{
	{
		nums:     []int{2, 7, 11, 15},
		target:   9,
		expected: []int{0, 1},
	},
	{
		nums:     []int{3, 2, 4},
		target:   6,
		expected: []int{1, 2},
	},
	{
		nums:     []int{3, 3},
		target:   6,
		expected: []int{0, 1},
	},
}
