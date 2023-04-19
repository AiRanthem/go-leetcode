package top100

func climbStairs(n int) int {
	now, one, two := 0, 1, 0 // i, i-1, i-2 阶楼梯的种数
	for i := 0; i < n; i++ {
		now = one + two
		one, two = now, one
	}
	return now
}
