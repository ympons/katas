package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxOperations(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, maxOperations(test.nums, test.k))
	}
}

func TestMaxOperationsMap(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, maxOperationsMap(test.nums, test.k))
	}
}

var tests = []struct {
	nums     []int
	k        int
	expected int
}{
	{
		nums:     []int{1, 2, 3, 4},
		k:        5,
		expected: 2,
	},
	{
		nums:     []int{3, 1, 3, 4, 3},
		k:        6,
		expected: 1,
	},
}
