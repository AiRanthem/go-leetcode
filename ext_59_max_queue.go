package leetcode

type MaxQueue struct {
	R  []int // Real
	M  []int // Max
	Ml int   // length
}

//func Constructor() MaxQueue {
//	return MaxQueue{
//		R:  []int{},
//		M:  []int{},
//		Ml: 0,
//	}
//}

func (q *MaxQueue) Max_value() int {
	if q.Ml == 0 {
		return -1
	}
	return q.M[0]
}

func (q *MaxQueue) Push_back(value int) {
	q.R = append(q.R, value)
	for q.Ml > 0 && value > q.M[q.Ml-1] {
		q.M = q.M[:q.Ml-1]
		q.Ml--
	}
	q.M = append(q.M, value)
	q.Ml++
}

func (q *MaxQueue) Pop_front() int {
	if q.Ml == 0 {
		return -1
	}
	val := q.R[0]
	q.R = q.R[1:]
	if val == q.Max_value() {
		q.M = q.M[1:]
		q.Ml--
	}
	return val
}
