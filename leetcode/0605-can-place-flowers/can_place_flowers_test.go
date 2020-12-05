package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := canPlaceFlowers(test.flowerbed, test.n)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	flowerbed []int
	n         int
	expected  bool
}{
	{
		flowerbed: []int{1, 0, 0, 0, 1},
		n:         1,
		expected:  true,
	},
	{
		flowerbed: []int{1, 0, 0, 0, 1},
		n:         2,
		expected:  false,
	},
	{
		flowerbed: []int{0, 0, 0},
		n:         2,
		expected:  true,
	},
	{
		flowerbed: []int{1, 0},
		n:         1,
		expected:  false,
	},
}
