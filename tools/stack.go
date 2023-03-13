package tools

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Top() int {
	return (*s)[s.Len()-1]
}

func (s *Stack) Pop() int {
	i := s.Top()
	*s = (*s)[:s.Len()-1]
	return i
}
