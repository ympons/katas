package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartitionKSubsets(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := canPartitionKSubsets(test.nums, test.k)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums     []int
	k        int
	expected bool
}{
	{
		nums:     []int{4, 3, 2, 3, 5, 2, 1},
		k:        4,
		expected: true,
	},
	{
		nums:     []int{10, 10, 10, 7, 7, 7, 7, 7, 7, 6, 6, 6},
		k:        3,
		expected: true,
	},
}
