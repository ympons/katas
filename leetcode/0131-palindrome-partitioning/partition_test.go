package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := partition(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    string
	expected [][]string
}{
	{
		input:    "aab",
		expected: [][]string{{"a", "a", "b"}, {"aa", "b"}},
	},
	{
		input:    "a",
		expected: [][]string{{"a"}},
	},
	{
		input:    "abc",
		expected: [][]string{{"a", "b", "c"}},
	},
}
