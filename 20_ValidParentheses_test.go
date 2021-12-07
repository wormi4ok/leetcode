package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test20_ValidParentheses(t *testing.T) {
	require.True(t, isValid("[()]{}"))
	require.True(t, isValid("[()]{}((((())))){}[[[{[{[]}]}]]]"))
	require.False(t, isValid("[()]}{}"))
	require.False(t, isValid("["))
	require.False(t, isValid(")"))
}

var pairs = map[rune]rune{
	'[': ']',
	'{': '}',
	'(': ')',
}

func match(open, close rune) bool {
	matching, exists := pairs[open]
	return exists && matching == close
}

type stack struct {
	elements []rune
}

func NewStack() *stack {
	return &stack{elements: make([]rune, 0, 4)}
}

func (s *stack) pop() {
	s.elements = s.elements[:len(s.elements)-1]
}

func (s *stack) push(el rune) {
	s.elements = append(s.elements, el)
}

func (s *stack) top() rune {
	return s.elements[len(s.elements)-1]
}

func (s *stack) isEmpty() bool {
	return len(s.elements) == 0
}

func isValid(s string) bool {
	st := NewStack()
	for _, char := range s {
		if !st.isEmpty() && match(st.top(), char) {
			st.pop()
		} else {
			st.push(char)
		}
	}
	return st.isEmpty()
}
