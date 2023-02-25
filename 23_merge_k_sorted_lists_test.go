package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	lists := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}},
		{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
	}
	head := mergeKLists(lists)
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}
