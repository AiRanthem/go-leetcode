package daily

import "container/heap"

type Heap1615 struct {
	H []int
	D []int               // degree
	L map[[2]int]struct{} // link
}

func (h *Heap1615) Len() int {
	return len(h.H)
}

func (h *Heap1615) Less(i, j int) bool {
	return h.D[h.H[i]] > h.D[h.H[j]]
}

func (h *Heap1615) Swap(i, j int) {
	h.H[i], h.H[j] = h.H[j], h.H[i]
}

func (h *Heap1615) Push(x any) {
	h.H = append(h.H, x.(int))
}

func (h *Heap1615) Pop() any {
	x := h.H[h.Len()-1]
	h.H = h.H[:h.Len()-1]
	return x
}

func (h *Heap1615) MaxD() int {
	return h.D[h.H[0]]
}

func (h *Heap1615) Link(i, j int) {
	h.L[[2]int{i, j}] = struct{}{}
	h.L[[2]int{j, i}] = struct{}{}
}

func (h *Heap1615) Linked(i, j int) bool {
	_, ok := h.L[[2]int{i, j}]
	return ok
}

func (h *Heap1615) Sum(i, j int) int {
	r := h.D[i] + h.D[j]
	if h.Linked(i, j) {
		r--
	}
	return r
}

func NewHeap1615(n int) *Heap1615 {
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = i
	}
	return &Heap1615{
		H: h,
		D: make([]int, n),
		L: map[[2]int]struct{}{},
	}
}

func maximalNetworkRank(n int, roads [][]int) int {
	if len(roads) == 0 {
		return 0
	}
	h := NewHeap1615(n)
	for i := 0; i < len(roads); i++ {
		h.D[roads[i][0]]++
		h.D[roads[i][1]]++
		h.Link(roads[i][0], roads[i][1])
	}
	heap.Init(h)
	var candidates []int
	for i := 0; i < 2 && h.Len() > 0; i++ {
		max := h.MaxD()
		for h.Len() > 0 && h.MaxD() == max {
			candidates = append(candidates, heap.Pop(h).(int))
		}
	}
	result := h.Sum(candidates[0], candidates[1])
	for i := 0; i < len(candidates); i++ {
		for j := i + 1; j < len(candidates); j++ {
			if r := h.Sum(candidates[i], candidates[j]); r > result {
				result = r
			}
		}
	}
	return result
}
