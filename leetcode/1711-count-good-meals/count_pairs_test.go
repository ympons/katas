package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPairs(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := countPairs(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{1, 3, 5, 7, 9},
		expected: 4,
	},
	{
		input:    []int{1, 1, 1, 3, 3, 3, 7},
		expected: 15,
	},
}
