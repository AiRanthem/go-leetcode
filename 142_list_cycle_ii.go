package leetcode

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			break
		}
	}
	return fast
}
