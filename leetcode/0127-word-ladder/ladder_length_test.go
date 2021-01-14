package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLadderLength(t *testing.T) {
	assert := assert.New(t)
	for _, test := range tests {
		assert.Equal(test.expected, ladderLength(test.beginWord, test.endWord, test.wordList))
	}
}

var tests = []struct {
	beginWord, endWord string
	wordList           []string
	expected           int
}{
	{
		beginWord: "leet",
		endWord:   "code",
		wordList:  []string{"lest", "leet", "lose", "code", "lode", "robe", "lost"},
		expected:  6,
	},
	{
		beginWord: "hit",
		endWord:   "cog",
		wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
		expected:  5,
	},
	{
		beginWord: "hit",
		endWord:   "cog",
		wordList:  []string{"hot", "dot", "dog", "lot", "log"},
		expected:  0,
	},
}
