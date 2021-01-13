package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEatenApples(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, eatenApples(test.apples, test.days))
	}
}

var tests = []struct {
	apples, days []int
	expected     int
}{
	{
		apples:   []int{1, 2, 3, 5, 2},
		days:     []int{3, 2, 1, 4, 2},
		expected: 7,
	},
	{
		apples:   []int{3, 0, 0, 0, 0, 2},
		days:     []int{3, 0, 0, 0, 0, 2},
		expected: 5,
	},
	{
		apples:   []int{2, 1, 10},
		days:     []int{2, 10, 1},
		expected: 4,
	},
}
