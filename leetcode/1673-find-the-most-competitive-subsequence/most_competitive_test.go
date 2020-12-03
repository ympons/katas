package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostCompetitive(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := mostCompetitive(test.nums, test.k)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums     []int
	k        int
	expected []int
}{
	{
		nums:     []int{3, 5, 2, 6},
		k:        2,
		expected: []int{2, 6},
	},
	{
		nums:     []int{2, 4, 3, 3, 5, 4, 9, 6},
		k:        4,
		expected: []int{2, 3, 3, 4},
	},
	{
		nums:     []int{71, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2},
		k:        3,
		expected: []int{8, 80, 2},
	},
}
