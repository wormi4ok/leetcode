package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test27_RemoveElement(t *testing.T) {
	input, val := []int{3, 2, 2, 3}, 3
	assert.Equal(t, 2, removeElement(input, val))
	assert.Equal(t, []int{2, 2, 0, 0}, input)

	input, val = []int{2, 2, 2, 3}, 3
	assert.Equal(t, 3, removeElement(input, val))
	assert.Equal(t, []int{2, 2, 2, 0}, input)

	input, val = []int{}, 1
	assert.Equal(t, 0, removeElement(input, val))
	assert.Equal(t, []int{}, input)

	input, val = []int{1, 1, 1}, 1
	assert.Equal(t, 0, removeElement(input, val))
	assert.Equal(t, []int{0, 0, 0}, input)

	input, val = []int{0, 1, 2, 2, 3, 0, 4, 2}, 2
	assert.Equal(t, 5, removeElement(input, val))
	assert.Equal(t, []int{0, 1, 3, 0, 4, 0, 0, 0}, input)
}

func removeElement(nums []int, val int) int {
	cnt := len(nums)
	gg := NewQueue()
	for i := range nums {
		if nums[i] == val {
			cnt--
			gg.add(i)
			nums[i] = 0
		} else if !gg.isEmpty() {
			nums[gg.remove()] = nums[i]
			gg.add(i)
			nums[i] = 0
		}
	}

	return cnt
}

type fifo struct {
	queue []int
}

func NewQueue() *fifo {
	return &fifo{queue: make([]int, 0, 4)}
}

func (s *fifo) add(el int) {
	s.queue = append(s.queue, el)
}

func (s *fifo) remove() int {
	first := s.queue[0]
	s.queue = s.queue[1:]
	return first
}

func (s *fifo) isEmpty() bool {
	return len(s.queue) == 0
}
