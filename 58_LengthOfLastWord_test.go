package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test58_LengthOfLastWord(t *testing.T) {
	assert.Equal(t, 5, lengthOfLastWord("Hello World"))
	assert.Equal(t, 4, lengthOfLastWord("   fly me   to   the moon  "))
	assert.Equal(t, 6, lengthOfLastWord("luffy is still joyboy"))
	assert.Equal(t, 1, lengthOfLastWord("a"))
}

func lengthOfLastWord(s string) int {
	var k int
	var isWord bool
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if isWord {
				return k
			}
			continue
		}
		isWord = true
		k++
	}
	return k
}
