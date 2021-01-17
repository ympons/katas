package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTupleSameProduct(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, tupleSameProduct(test.input))
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{2, 3, 4, 6},
		expected: 8,
	},
	{
		input:    []int{1, 2, 4, 5, 10},
		expected: 16,
	},
	{
		input:    []int{2, 3, 4, 6, 8, 12},
		expected: 40,
	},
	{
		input:    []int{2, 3, 5, 7},
		expected: 0,
	},
}
