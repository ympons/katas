package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoneGameVI(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		output := stoneGameVI(test.aliceValues, test.bobValues)
		assert.Equal(test.expected, output)
	}
}

var tests = []struct {
	aliceValues, bobValues []int
	expected               int
}{
	{
		aliceValues: []int{1, 3},
		bobValues:   []int{2, 1},
		expected:    1,
	},
	{
		aliceValues: []int{1, 2},
		bobValues:   []int{3, 1},
		expected:    0,
	},
	{
		aliceValues: []int{2, 4, 3},
		bobValues:   []int{1, 6, 7},
		expected:    -1,
	},
}
