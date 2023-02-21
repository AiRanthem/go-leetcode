package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	for {
		if list1 == nil {
			cur.Next = list2
			return root.Next
		}
		if list2 == nil {
			cur.Next = list1
			return root.Next
		}
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
}
