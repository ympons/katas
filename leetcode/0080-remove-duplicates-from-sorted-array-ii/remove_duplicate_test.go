package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicate(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := removeDuplicate(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{1, 1, 1, 2, 2, 3},
		expected: 5,
	},
	{
		input:    []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
		expected: 7,
	},
}
