package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test9_PalindromeNumber(t *testing.T) {
	require.True(t, isPalindrome(123321))
	require.True(t, isPalindrome(1237321))
	require.False(t, isPalindrome(12237321))
	require.False(t, isPalindrome(-11))
	require.True(t, isPalindrome(11))
	require.True(t, isPalindrome(1))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := fmt.Sprintf("%d", x)
	for head := range num {
		tail := len(num) - head
		if num[head] != num[tail-1] {
			return false
		}
	}

	return true
}
