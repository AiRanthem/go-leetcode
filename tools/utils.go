package tools

import "strconv"

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

func minTInt(x, y, z int) int {
	return minInt(minInt(x, y), z)
}

func atoi(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func sum(arr []int) (s int) {
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return
}
