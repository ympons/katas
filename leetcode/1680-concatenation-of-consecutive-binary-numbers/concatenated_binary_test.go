package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatenatedBinary(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := concatenatedBinary(test.n)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	n        int
	expected int
}{
	{
		n:        1,
		expected: 1,
	},
	{
		n:        3,
		expected: 27,
	},
	{
		n:        12,
		expected: 505379714,
	},
}
