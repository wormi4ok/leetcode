package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test26_RemoveDuplicatesFromSortedArray(t *testing.T) {

	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	want := []int{0, 1, 2, 3, 4, 0, 0, 0, 0, 0}

	assert.Equal(t, 5, removeDuplicates(input))
	assert.Equal(t, want, input)

	input = []int{1, 1, 2, 3, 3}
	want = []int{1, 2, 3, 0, 0}

	assert.Equal(t, 3, removeDuplicates(input))
	assert.Equal(t, want, input)

	input = []int{1, 1, 2}
	want = []int{1, 2, 0}

	assert.Equal(t, 2, removeDuplicates(input))
	assert.Equal(t, want, input)

	input = []int{1, 2, 3}
	want = []int{1, 2, 3}

	assert.Equal(t, 3, removeDuplicates(input))
	assert.Equal(t, want, input)

	input = []int{}
	want = []int{}

	assert.Equal(t, 0, removeDuplicates(input))
	assert.Equal(t, want, input)
}

func removeDuplicates(nums []int) int {
	j := 0
	length := len(nums)

	if length == 0 {
		return 0
	}

	for i := 1; i < length; i++ {
		if nums[j] != nums[i] {
			j++
		}
		nums[j] = nums[i]

		if i != j {
			nums[i] = 0
		}
	}

	return j + 1
}
