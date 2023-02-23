package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	s := Stack{}
	for _, token := range tokens {
		switch token {
		case "+":
			b, a := s.Pop(), s.Pop()
			s.Push(a + b)
		case "-":
			b, a := s.Pop(), s.Pop()
			s.Push(a - b)
		case "*":
			b, a := s.Pop(), s.Pop()
			s.Push(a * b)
		case "/":
			b, a := s.Pop(), s.Pop()
			s.Push(a / b)
		default:
			i, _ := strconv.Atoi(token)
			s.Push(i)
		}
	}
	return s.Top()
}
