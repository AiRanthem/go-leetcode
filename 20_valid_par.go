package leetcode

type ParenthesStack []byte

func (s *ParenthesStack) Push(p byte) {
	*s = append(*s, p)
}

func (s *ParenthesStack) Pop() (val byte) {
	l := len(*s)
	if l == 0 {
		return 0
	}
	val = (*s)[l-1]
	*s = (*s)[:l-1]
	return
}

func isValid(s string) bool {
	ss := []byte(s)
	st := &ParenthesStack{}
	for i := 0; i < len(ss); i++ {
		switch ss[i] {
		case '(':
			st.Push(ss[i])
		case '[':
			st.Push(ss[i])
		case '{':
			st.Push(ss[i])
		case ')':
			if val := st.Pop(); val != '(' {
				return false
			}
		case ']':
			if val := st.Pop(); val != '[' {
				return false
			}
		default:
			if val := st.Pop(); val != '{' {
				return false
			}
		}
	}
	return len(*st) == 0
}
