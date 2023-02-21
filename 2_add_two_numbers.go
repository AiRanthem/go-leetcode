package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	var x, y, z, s int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			cur.Next = l1
			x = l1.Val
		} else {
			x = 0
		}
		if l2 != nil {
			cur.Next = l2
			y = l2.Val
		} else {
			y = 0
		}
		cur = cur.Next
		if s = x + y + z; s >= 10 {
			s -= 10
			z = 1
		} else {
			z = 0
		}
		cur.Val = s
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if z != 0 {
		cur.Next = &ListNode{Val: z}
	}
	return root.Next
}
