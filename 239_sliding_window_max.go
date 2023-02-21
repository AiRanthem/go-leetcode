package leetcode

// 单调队列

type MonoQueue []int

func (q *MonoQueue) Front() int {
	return (*q)[0]
}

func (q *MonoQueue) Tail() int {
	return (*q)[len(*q)-1]
}

func (q *MonoQueue) Push(x int) {
	for len(*q) > 0 && x > q.Tail() {
		*q = (*q)[:len(*q)-1]
	}
	*q = append(*q, x)
}

func (q *MonoQueue) Pop(x int) {
	if x == q.Front() {
		*q = (*q)[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	q := &MonoQueue{}
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	n := len(nums)
	result := make([]int, n-k+1)
	for i := k; i < n; i++ {
		result[i-k] = q.Front()
		q.Push(nums[i])
		q.Pop(nums[i-k])
	}
	result[n-k] = q.Front()
	return result
}
