package leetcode

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test40_CombinationSum2(t *testing.T) {
	assert.Equal(t, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	assert.Equal(t, [][]int{{5, 5, 5}, {7, 8}}, combinationSum2([]int{5, 6, 5, 7, 5, 8}, 15))
	assert.Equal(t, [][]int{{3, 4}}, combinationSum2([]int{3, 4, 8, 9, 10}, 7))
	assert.Equal(t, [][]int{{1, 2, 2}, {5}}, combinationSum2([]int{2, 5, 2, 1, 2}, 5))
	assert.Equal(t, [][]int{{2, 2}}, combinationSum2([]int{2, 2, 2}, 4))
	assert.Equal(t, [][]int{}, combinationSum2([]int{1}, 2))
}

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	slices.Sort(candidates)
	// optional optimisation to exclude values higher than target
	candidates = reduceCandidates(candidates, target)

	findSum(&result, candidates, []int{}, 0, target)
	return result
}

func findSum(result *[][]int, candidates []int, current []int, position int, target int) {
	if target == 0 {
		// copy current combination to avoid using referenced value
		res := make([]int, len(current))
		copy(res, current)
		*result = append(*result, res)
		return
	}
	if target < 0 {
		return
	}
	prev := -1
	for i := position; i < len(candidates); i++ {
		candidate := candidates[i]
		if candidate == prev {
			continue
		}
		current = append(current, candidate)
		findSum(result, candidates, current, i+1, target-candidate)
		current = current[:len(current)-1]
		prev = candidate
	}
}

func reduceCandidates(candidates []int, target int) []int {
	for i := len(candidates) - 1; i > 0; i-- {
		if candidates[i] <= target {
			return candidates
		}
		candidates = candidates[:i]
	}
	return candidates
}
