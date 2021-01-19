package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSortedArray(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, createSortedArray(test.input))
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{1, 5, 6, 2},
		expected: 1,
	},
	{
		input:    []int{1, 2, 3, 6, 5, 4},
		expected: 3,
	},
	{
		input:    []int{1, 3, 3, 3, 2, 4, 2, 1, 2},
		expected: 4,
	},
}
