package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumDeviation(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, minimumDeviation(test.input))
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{1, 2, 3, 4},
		expected: 1,
	},
	{
		input:    []int{4, 1, 5, 20, 3},
		expected: 3,
	},
	{
		input:    []int{2, 10, 8},
		expected: 3,
	},
}
