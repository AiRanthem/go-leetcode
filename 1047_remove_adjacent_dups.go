package leetcode

import "strings"

func removeDuplicates(s string) string {
	stack := Stack{}
	for i := 0; i < len(s); i++ {
		if c := int(s[i]); stack.Len() > 0 && stack.Top() == c {
			stack.Pop()
		} else {
			stack.Push(c)
		}
	}
	sb := strings.Builder{}
	for i := 0; i < stack.Len(); i++ {
		sb.WriteByte(byte(stack[i]))
	}
	return sb.String()
}
