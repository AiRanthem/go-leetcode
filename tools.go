package leetcode

import "strconv"

func new2dSlice(m, n int) [][]int {
	s := make([][]int, m)
	for i := 0; i < m; i++ {
		s[i] = make([]int, n)
	}
	return s
}

func maxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func minInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func minTInt(x, y, z int) int {
	return minInt(minInt(x, y), z)
}

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

type ChildNode struct {
	Val      int
	Children []*ChildNode
}

func atoi(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}
