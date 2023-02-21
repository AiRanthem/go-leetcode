package leetcode

func swapPairs(head *ListNode) *ListNode {
	cur := head
	if head != nil && head.Next != nil {
		head = head.Next
	}
	var next, tmp, pre *ListNode
	for cur != nil && cur.Next != nil {
		next = cur.Next
		tmp = next.Next
		next.Next = cur
		cur.Next = tmp
		if pre != nil {
			pre.Next = next
		}
		pre = cur
		cur = tmp
	}
	return head
}
