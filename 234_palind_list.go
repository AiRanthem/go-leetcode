package leetcode

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	var next, before *ListNode
	next = slow.Next
	var even bool
	for fast.Next != nil {
		fast = fast.Next
		if fast.Next == nil {
			even = true
			break
		}
		fast = fast.Next
		slow.Next = before
		before = slow
		slow = next
		next = slow.Next
	}
	if even {
		slow.Next = before
	} else {
		slow = before
	}
	for next != nil || slow != nil {
		if slow == nil || next == nil || slow.Val != next.Val {
			return false
		}
		slow = slow.Next
		next = next.Next
	}
	return true
}
