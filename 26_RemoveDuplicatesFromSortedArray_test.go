package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test26_RemoveDuplicatesFromSortedArray(t *testing.T) {

	tests := []struct {
		input  []int
		result []int
		length int
	}{{
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		[]int{0, 1, 2, 3, 4, 0, 0, 0, 0, 0},
		5,
	}, {
		[]int{1, 1, 2, 3, 3},
		[]int{1, 2, 3, 0, 0},
		3,
	}, {
		[]int{1, 1, 2},
		[]int{1, 2, 0},
		2,
	}, {
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		3,
	}, {
		[]int{},
		[]int{},
		0,
	},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.length, removeDuplicates(tt.input))
		assert.Equal(t, tt.result, tt.input)
	}
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
