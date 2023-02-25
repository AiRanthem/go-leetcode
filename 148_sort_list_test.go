package leetcode

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {
	l := sortList(&ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}})
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
}
