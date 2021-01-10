package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumGain(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, maximumGain(test.s, test.x, test.y))
	}
}

var tests = []struct {
	s        string
	x, y     int
	expected int
}{
	{
		s:        "cdbcbbaaabab",
		x:        4,
		y:        5,
		expected: 19,
	},
	{
		s:        "aabbaaxybbaabb",
		x:        5,
		y:        4,
		expected: 20,
	},
}
