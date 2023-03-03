package leetcode

type TopKHeap struct {
	b [][2]int // [number, frequency]
	k int
	l int
}

func NewTopKHeap(k int) *TopKHeap {
	return &TopKHeap{
		b: make([][2]int, k),
		k: k,
		l: 0,
	}
}

func (h *TopKHeap) up(i int) {
	p := (i - 1) / 2
	for p >= 0 && h.b[i][1] < h.b[p][1] {
		h.b[p], h.b[i] = h.b[i], h.b[p]
		i = p
		p = (i - 1) / 2
	}
}

func (h *TopKHeap) down(i int) {
	c := 2*i + 1
	for c < h.l {
		if c < h.l-1 && h.b[c+1][1] < h.b[c][1] {
			c++
		}
		if h.b[i][1] > h.b[c][1] {
			h.b[c], h.b[i] = h.b[i], h.b[c]
			i = c
			c = 2*i + 1
		} else {
			break
		}
	}
}

func (h *TopKHeap) Push(num int, freq int) {
	if h.l == h.k {
		if freq < h.b[0][1] {
			return
		}
		h.Pop()
	}
	h.b[h.l] = [2]int{num, freq}
	h.up(h.l)
	h.l++
}

func (h *TopKHeap) Pop() (elem [2]int) {
	elem = h.b[0]
	h.b[0], h.b[h.l-1] = h.b[h.l-1], h.b[0]
	h.l--
	h.down(0)
	return
}

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}
	h := NewTopKHeap(k)
	for num, freq := range counter {
		h.Push(num, freq)
	}
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = h.Pop()[0]
	}
	return result
}
