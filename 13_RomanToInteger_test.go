package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test13_RomanToInteger(t *testing.T) {
	require.Equal(t, 3, romanToInt("III"))
	require.Equal(t, 4, romanToInt("IV"))
	require.Equal(t, 9, romanToInt("IX"))
	require.Equal(t, 58, romanToInt("LVIII"))
	require.Equal(t, 1994, romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	sum := 0
	rng := len(s) - 1

	for i := range s {
		sub := 0
		if i < rng {
			sub = subtract(s[i], s[i+1])
		}
		if sub > 0 {
			sum = sum - sub
			continue
		}
		sum = sum + rn[rune(s[i])]
	}
	return sum
}

var rn = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func subtract(letter, nextLetter uint8) int {
	switch letter {
	case 'I':
		if nextLetter == 'V' || nextLetter == 'X' {
			return 1
		}
	case 'X':
		if nextLetter == 'L' || nextLetter == 'C' {
			return 10
		}
	case 'C':
		if nextLetter == 'D' || nextLetter == 'M' {
			return 100
		}
	}

	return 0
}
