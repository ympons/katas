package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPredictTheWinner(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := PredictTheWinner(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    []int
	expected bool
}{
	{
		input:    []int{1, 5, 2},
		expected: false,
	},
	{
		input:    []int{1, 5, 233, 7},
		expected: true,
	},
	{
		input:    []int{1, 2},
		expected: true,
	},
	{
		input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		expected: true,
	},
}
