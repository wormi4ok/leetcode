package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test35_SearchInsertPosition(t *testing.T) {
	assert.Equal(t, 1, searchInsert([]int{2, 4, 6, 8}, 3))
	assert.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.Equal(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))
	assert.Equal(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))
}

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		m := (i + j) >> 1
		if nums[m] >= target {
			j = m
		} else {
			i = m + 1
		}
	}

	return i
}
