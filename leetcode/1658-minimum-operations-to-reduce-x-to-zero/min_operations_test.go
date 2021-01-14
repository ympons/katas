package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, minOperations(test.nums, test.x))
	}
}

func TestMinOperationsMap(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, minOperationsMap(test.nums, test.x))
	}
}

var tests = []struct {
	nums     []int
	x        int
	expected int
}{
	{
		nums:     []int{1, 1, 4, 2, 3},
		x:        5,
		expected: 2,
	},
	{
		nums:     []int{5, 6, 7, 8, 9},
		x:        4,
		expected: -1,
	},
	{
		nums:     []int{3, 2, 20, 1, 1, 3},
		x:        10,
		expected: 5,
	},
}
