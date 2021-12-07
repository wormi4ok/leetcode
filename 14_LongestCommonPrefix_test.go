package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test14_LongestCommonPrefix(t *testing.T) {
	require.Equal(t, "abc", longestCommonPrefix([]string{"abcde", "abczyx", "abc555"}))
	require.Equal(t, "1", longestCommonPrefix([]string{"123abcde", "12", "11111"}))
	require.Equal(t, "", longestCommonPrefix([]string{"123abcde", "xyz"}))
	require.Equal(t, "", longestCommonPrefix([]string{""}))
	require.Equal(t, "", longestCommonPrefix([]string{"333", "", "333"}))
	require.Equal(t, "a", longestCommonPrefix([]string{"ab", "a"}))
}

func longestCommonPrefix(strs []string) string {
	prefix := ""
	for i, str := range strs {
		length := len(str)
		if i == 0 || length == 0 {
			prefix = str
			continue
		}
		if len(prefix) > length {
			prefix = prefix[:length]
		}
		for j := range str {
			if len(prefix) > j && prefix[j] != str[j] {
				prefix = str[:j]
				break
			}
		}
	}

	return prefix
}
