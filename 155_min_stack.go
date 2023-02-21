package leetcode

type MinStack struct {
	S []int // real stack
	M []int // min  stack
}

//func Constructor() MinStack {
//	return MinStack{[]int{}, []int{math.MaxInt}}
//}

func (m *MinStack) Push(val int) {
	m.S = append(m.S, val)
	mIdx := len(m.M) - 1
	if val <= m.M[mIdx] {
		m.M = append(m.M, val)
	}
}

func (m *MinStack) Pop() {
	sIdx := len(m.S) - 1
	mIdx := len(m.M) - 1
	val := m.S[sIdx]
	m.S = m.S[:sIdx]
	if val == m.M[mIdx] {
		m.M = m.M[:mIdx]
	}
}

func (m *MinStack) Top() int {
	return m.S[len(m.S)-1]
}

func (m *MinStack) GetMin() int {
	return m.M[len(m.M)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushRight(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
