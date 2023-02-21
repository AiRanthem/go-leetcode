package leetcode

func reverseList(head *ListNode) *ListNode {
	var pre, cur, tmp *ListNode
	cur = head
	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
