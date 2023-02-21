package leetcode

import "container/heap"

type IHeap [][]int // [number, freq]

func (h *IHeap) Len() int {
	return len(*h)
}

func (h *IHeap) Less(i, j int) bool {
	return (*h)[i][1] < (*h)[j][1]
}

func (h *IHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	*h = old[:n-1]
	return old[n-1]
}

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	h := &IHeap{}
	heap.Init(h)

	for n, freq := range counter {
		heap.Push(h, []int{n, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).([]int)[0]
	}
	return result
}
