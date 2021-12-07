package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1_TwoSum(t *testing.T) {
	require.ElementsMatch(t, []int{1, 2}, twoSum([]int{1, 2, 4}, 6))
	require.ElementsMatch(t, []int{0, 2}, twoSum([]int{12, 2, 5, 7, 8}, 17))
	require.ElementsMatch(t, []int{3, 4}, twoSum([]int{12, 2, 5, 7, 8}, 15))
}

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))

	for i, num := range nums {
		complement := target - num
		if j, found := hashmap[complement]; found {
			return []int{i, j}
		}
		hashmap[num] = i
	}

	return nil
}
