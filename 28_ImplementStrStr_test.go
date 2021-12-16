package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test28_ImplementStrStr(t *testing.T) {
	assert.Equal(t, 2, strStr("hello", "ll"))
	assert.Equal(t, -1, strStr("aaaaa", "bba"))
	assert.Equal(t, 0, strStr("", ""))
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
