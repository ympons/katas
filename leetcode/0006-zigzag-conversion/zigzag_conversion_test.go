package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := convert(test.s, test.numRows)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	s        string
	numRows  int
	expected string
}{
	{
		s:        "PAYPALISHIRING",
		numRows:  3,
		expected: "PAHNAPLSIIGYIR",
	},
	{
		s:        "PAYPALISHIRING",
		numRows:  4,
		expected: "PINALSIGYAHRPI",
	},
	{
		s:        "A",
		numRows:  1,
		expected: "A",
	},
	{
		s:        "PAYPALISHIRING",
		numRows:  5,
		expected: "PHASIYIRPLIGAN",
	},
	{
		s:        "PAYPALISHIRING",
		numRows:  6,
		expected: "PRAIIYHNPSGAIL",
	},
	{
		s:        "PAYPALISHIRING",
		numRows:  2,
		expected: "PYAIHRNAPLSIIG",
	},
}
