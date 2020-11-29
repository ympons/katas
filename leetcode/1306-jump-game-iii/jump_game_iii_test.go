package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanReach(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := canReach(test.nums, test.start)
		assert.Equal(test.expected, output)
	}
}

func TestCanReachDFS(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := canReachDFS(test.nums, test.start)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums     []int
	start    int
	expected bool
}{
	{
		nums:     []int{4, 2, 3, 0, 3, 1, 2},
		start:    5,
		expected: true,
	},
	{
		nums:     []int{4, 2, 3, 0, 3, 1, 2},
		start:    0,
		expected: true,
	},
	{
		nums:     []int{3, 0, 2, 1, 2},
		start:    2,
		expected: false,
	},
	{
		nums:     []int{1, 1, 1, 1, 1, 1, 1, 1, 0},
		start:    3,
		expected: true,
	},
}
