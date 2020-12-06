package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackspaceCompare(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := backspaceCompare(test.S, test.T)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	S, T     string
	expected bool
}{
	{
		S:        "ab#c",
		T:        "ad#c",
		expected: true,
	},
	{
		S:        "ab##",
		T:        "c#d#",
		expected: true,
	},
	{
		S:        "a##c",
		T:        "#a#c",
		expected: true,
	},
	{
		S:        "a#c",
		T:        "b",
		expected: false,
	},
}
