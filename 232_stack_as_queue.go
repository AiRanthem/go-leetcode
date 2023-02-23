package leetcode

type StackToBeQueue []int

func (s *StackToBeQueue) Push(x int) {
	*s = append(*s, x)
}

func (s *StackToBeQueue) Pop() int {
	i := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return i
}

func (s *StackToBeQueue) Peak() int {
	return (*s)[len(*s)-1]
}

type MyQueue struct {
	W *StackToBeQueue
	R *StackToBeQueue
}

//func Constructor() MyQueue {
//	return MyQueue{&StackToBeQueue{}, &StackToBeQueue{}}
//}

func (q *MyQueue) Push(x int) {
	q.W.Push(x)
}

func (q *MyQueue) Pop() int {
	q.refresh()
	return q.R.Pop()
}

func (q *MyQueue) Peek() int {
	q.refresh()
	return q.R.Peak()
}

func (q *MyQueue) Empty() bool {
	return len(*(q.R)) == 0 && len(*(q.W)) == 0
}

func (q *MyQueue) refresh() {
	if len(*(q.R)) == 0 {
		for len(*(q.W)) > 0 {
			q.R.Push(q.W.Pop())
		}
	}
}
