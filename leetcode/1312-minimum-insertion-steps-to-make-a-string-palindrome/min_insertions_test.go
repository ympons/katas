package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinInsertions(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minInsertions(test.input)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	input    string
	expected int
}{
	{
		input:    "zzazz",
		expected: 0,
	},
	{
		input:    "mbadm",
		expected: 2,
	},
	{
		input:    "leetcode",
		expected: 5,
	},
	{
		input:    "g",
		expected: 0,
	},
	{
		input:    "no",
		expected: 1,
	},
}
