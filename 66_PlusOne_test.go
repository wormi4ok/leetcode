package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test66_PlusOne(t *testing.T) {
	assert.Equal(t, []int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	assert.Equal(t, []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
	assert.Equal(t, []int{1, 0}, plusOne([]int{9}))
	assert.Equal(t, []int{1, 0, 0, 0, 0}, plusOne([]int{9, 9, 9, 9}))
	assert.Equal(t, []int{9, 3, 0}, plusOne([]int{9, 2, 9}))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
