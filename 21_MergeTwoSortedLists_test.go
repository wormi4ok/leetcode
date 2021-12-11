package leetcode

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test21_MergeTwoSortedLists(t *testing.T) {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	want := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}

	validate(t, want, mergeTwoLists(list1, list2))

	list1 = nil
	list2 = nil
	want = nil

	validate(t, want, mergeTwoLists(list1, list2))

	list1 = nil
	list2 = &ListNode{0, nil}
	want = &ListNode{0, nil}

	validate(t, want, mergeTwoLists(list1, list2))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func validate(t *testing.T, expected *ListNode, actual *ListNode) {
	for i := 0; i <= 100; i++ {
		if expected == nil && actual == nil {
			return
		}

		if expected != nil && actual == nil || expected == nil && actual != nil {
			t.Fatalf("Value #%d is missing in one of the lists", i)
		}

		if expected != nil && actual != nil && expected.Val != actual.Val {
			t.Fatalf("Value #%d in the list item is: %d; expected: %d", i, actual.Val, expected.Val)
		}

		expected = expected.Next
		actual = actual.Next
	}
}
