package leetcode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	a, b := splitList(head)
	root := &ListNode{}
	cur := root
	for a != nil || b != nil {
		if a == nil {
			cur.Next = b
			break
		}
		if b == nil {
			cur.Next = a
			break
		}
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}
	return root.Next
}

func splitList(head *ListNode) (*ListNode, *ListNode) {
	slow, fast := head, head
	for {
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		if fast == nil {
			break
		}
		slow = slow.Next
	}
	fast = slow.Next
	slow.Next = nil
	return sortList(head), sortList(fast)
}
