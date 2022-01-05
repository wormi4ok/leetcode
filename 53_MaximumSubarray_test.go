package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test53_MaximumSubarray(t *testing.T) {
	assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, 1, maxSubArray([]int{1}))
	assert.Equal(t, 23, maxSubArray([]int{5, 4, -1, 7, 8}))
	assert.Equal(t, 12, maxSubArray([]int{11, -2, -3, -4, 10}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var maxSum, currentSum = nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(currentSum+nums[i], nums[i])
		maxSum = max(currentSum, maxSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
