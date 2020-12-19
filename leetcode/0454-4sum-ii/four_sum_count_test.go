package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSumCount(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := fourSumCount(test.A, test.B, test.C, test.D)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	A, B, C, D []int
	expected   int
}{
	{
		A:        []int{1, 2},
		B:        []int{-2, -1},
		C:        []int{-1, 2},
		D:        []int{0, 2},
		expected: 2,
	},
}
