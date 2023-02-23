package leetcode

func reverseList(head *ListNode) *ListNode {
	var pre, cur, next *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
