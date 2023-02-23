package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	root := &ListNode{Next: head}
	pre, slow, fast := root, root, root
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
