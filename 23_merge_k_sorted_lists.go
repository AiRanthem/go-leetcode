package leetcode

import "container/heap"

type MKLHeap []*ListNode

func (h *MKLHeap) Len() int {
	return len(*h)
}

func (h *MKLHeap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}

func (h *MKLHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MKLHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *MKLHeap) Pop() any {
	i := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return i
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := MKLHeap{}
	heap.Init(&h)
	root := &ListNode{}
	cur := root
	for i, list := range lists {
		if list != nil {
			heap.Push(&h, list)
			lists[i] = list.Next
		}
	}
	for h.Len() > 0 {
		node := heap.Pop(&h).(*ListNode)
		cur.Next = node
		cur = cur.Next
		for i, list := range lists {
			if list != nil {
				heap.Push(&h, list)
				lists[i] = list.Next
			}
		}
	}
	cur.Next = nil
	return root.Next
}
