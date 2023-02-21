package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	root := &ListNode{0, head}
	cur, next := root, head
	for next != nil {
		if next.Val != val {
			cur = next
			next = cur.Next
		} else {
			cur.Next = next.Next
			next.Next = nil
			next = cur.Next
		}
	}
	return root.Next
}
