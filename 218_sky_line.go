package leetcode

import (
	"container/heap"
	"sort"
)

type SkyLineHeap [][]int // [left, right, height]

func (s *SkyLineHeap) Len() int {
	return len(*s)
}

func (s *SkyLineHeap) Less(i, j int) bool {
	return (*s)[i][2] > (*s)[j][2]
}

func (s *SkyLineHeap) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *SkyLineHeap) Push(x any) {
	*s = append(*s, x.([]int))
}

func (s *SkyLineHeap) Pop() any {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func getSkyline(buildings [][]int) (result [][]int) {
	points := make([]int, len(buildings)*2)
	for i := 0; i < len(buildings); i++ {
		points[2*i] = buildings[i][0]
		points[2*i+1] = buildings[i][1]
	}
	sort.Ints(points)
	h := SkyLineHeap{[]int{0, points[len(points)-1] + 1, 0}}
	bc := 0         // building counter
	var highest int // a helper buffer to record highest height at some scanned point
	for i := 0; i < len(points); i++ {
		// handle left bound
		highest = h[0][2]
		for bc < len(buildings) && buildings[bc][0] == points[i] {
			heap.Push(&h, buildings[bc])
			bc++
		}
		if highest < h[0][2] {
			result = append(result, []int{points[i], h[0][2]})
		}
		// hancle right bound. a build's height contributes to the skyline only when its right > scanned point.
		highest = h[0][2]
		for h[0][1] <= points[i] {
			// delete all useless buildings from the tallest
			// the first time this loop executed is when the scanned point is the right bound of currently the tallest building
			heap.Pop(&h)
		}
		if highest > h[0][2] {
			result = append(result, []int{points[i], h[0][2]})
		}
	}
	return
}
