package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minOperations(test.target, test.arr)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	target, arr []int
	expected    int
}{
	{
		target:   []int{5, 1, 3},
		arr:      []int{9, 4, 2, 3, 4},
		expected: 2,
	},
	{
		target:   []int{6, 4, 8, 1, 3, 2},
		arr:      []int{4, 7, 6, 2, 3, 8, 6, 1},
		expected: 3,
	},
}
