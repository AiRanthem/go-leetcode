package leetcode

func new2dSlice(m, n int) [][]int {
	s := make([][]int, m)
	for i := 0; i < m; i++ {
		s[i] = make([]int, n)
	}
	return s
}

func maxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func minInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
