package top100

func climbStairs(n int) int {
	now, one, two := 1, 1, 1
	for i := 0; i < n-1; i++ {
		now = one + two
		two = one
		one = now
	}
	return now
}
