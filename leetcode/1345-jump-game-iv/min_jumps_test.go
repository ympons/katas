package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinJumps(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := minJumps(test.nums)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	nums     []int
	expected int
}{
	{
		nums:     []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404},
		expected: 3,
	},
	{
		nums:     []int{7},
		expected: 0,
	},
	{
		nums:     []int{7, 6, 9, 6, 9, 6, 9, 7},
		expected: 1,
	},
	{
		nums:     []int{6, 1, 9},
		expected: 2,
	},
	{
		nums:     []int{11, 22, 7, 7, 7, 7, 7, 7, 7, 22, 13},
		expected: 3,
	},
	{
		nums:     []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 11},
		expected: 2,
	},
}
