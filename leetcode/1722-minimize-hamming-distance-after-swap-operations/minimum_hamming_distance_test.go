package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumHammingDistance(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minimumHammingDistance(test.source, test.target, test.allowedSwaps)
		assert.Equal(test.expected, output)
	}
}

func TestMinimumHammingDistanceUF(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minimumHammingDistanceUF(test.source, test.target, test.allowedSwaps)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	source, target []int
	allowedSwaps   [][]int
	expected       int
}{
	{
		source:       []int{1, 2, 3, 4},
		target:       []int{2, 1, 4, 5},
		allowedSwaps: [][]int{{0, 1}, {2, 3}},
		expected:     1,
	},
	{
		source:       []int{1, 2, 3, 4},
		target:       []int{1, 3, 2, 4},
		allowedSwaps: [][]int{},
		expected:     2,
	},
	{
		source:       []int{5, 1, 2, 4, 3},
		target:       []int{1, 5, 4, 2, 3},
		allowedSwaps: [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}},
		expected:     0,
	},
}
