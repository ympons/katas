package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := uniquePaths(test.m, test.n)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	m, n     int
	expected int
}{
	{
		m:        3,
		n:        7,
		expected: 28,
	},
	{
		m:        3,
		n:        2,
		expected: 3,
	},
	{
		m:        7,
		n:        3,
		expected: 28,
	},
	{
		m:        3,
		n:        3,
		expected: 6,
	},
}
