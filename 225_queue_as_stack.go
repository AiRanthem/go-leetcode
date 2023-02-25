package leetcode

type QueueToBeStack []int

func (q *QueueToBeStack) Push(x int) {
	*q = append(*q, x)
}

func (q *QueueToBeStack) Pop() int {
	i := (*q)[0]
	*q = (*q)[1:]
	return i
}

func (q *QueueToBeStack) Peek() int {
	return (*q)[0]
}

type MyStack struct {
	R *QueueToBeStack
	W *QueueToBeStack
}

func (s *MyStack) refresh() {
	for len(*(s.W)) > 1 {
		s.R.Push(s.W.Pop())
	}
}

//func Constructor() MyStack {
//	return MyStack{&QueueToBeStack{}, &QueueToBeStack{}}
//}

func (s *MyStack) Push(x int) {
	s.W.Push(x)
}

func (s *MyStack) Pop() int {
	s.refresh()
	i := s.W.Pop()
	s.W, s.R = s.R, s.W
	return i
}

func (s *MyStack) Top() int {
	s.refresh()
	return s.W.Peek()
}

func (s *MyStack) Empty() bool {
	return len(*(s.R)) == 0 && len(*(s.W)) == 0
}
