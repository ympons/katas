package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		ouput := calculate(test.input)
		assert.Equal(test.expected, ouput)
	}
}

var tests = []struct {
	input    string
	expected int
}{
	{
		input:    "3+2*2",
		expected: 7,
	},
	{
		input:    " 3/2 ",
		expected: 1,
	},
	{
		input:    " 3+5 / 2 ",
		expected: 5,
	},
	{
		input:    "14-3/2",
		expected: 13,
	},
}
