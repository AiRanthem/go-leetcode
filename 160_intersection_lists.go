package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	la, lb := 0, 0
	ha, hb := headA, headB
	for ha != nil {
		la++
		ha = ha.Next
	}
	for hb != nil {
		lb++
		hb = hb.Next
	}
	if la < lb {
		headA, headB = headB, headA
		la, lb = lb, la
	}
	for i := 0; i < la-lb; i++ {
		headA = headA.Next
	}
	for headA != nil {
		if headA == headB {
			break
		}
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
