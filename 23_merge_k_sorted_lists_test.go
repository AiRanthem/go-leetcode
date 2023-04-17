package leetcode

import "testing"

func TestMergeKLists(t *testing.T) {
	mergeKLists([]*ListNode{
		{1, &ListNode{4, &ListNode{5, nil}}},
		{1, &ListNode{3, &ListNode{4, nil}}},
		{2, &ListNode{6, nil}},
	})
}
