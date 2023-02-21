package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	root := &ListNode{0, head}
	slow, fast := root, root
	var pre *ListNode
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	pre.Next = slow.Next
	return root.Next
}
