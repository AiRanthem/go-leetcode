package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var la, lb int
	var cur *ListNode
	cur = headA
	for cur != nil {
		la++
		cur = cur.Next
	}
	cur = headB
	for cur != nil {
		lb++
		cur = cur.Next
	}
	if la < lb {
		la, lb = lb, la
		headA, headB = headB, headA
	}
	for i := 0; i < la-lb; i++ {
		headA = headA.Next
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
