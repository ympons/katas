package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructDistancedSequence(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := constructDistancedSequence(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    int
	expected []int
}{
	{
		input:    1,
		expected: []int{1},
	},
	{
		input:    3,
		expected: []int{3, 1, 2, 3, 2},
	},
	{
		input:    5,
		expected: []int{5, 3, 1, 4, 3, 5, 2, 4, 2},
	},
}
