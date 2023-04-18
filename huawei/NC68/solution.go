package main

// 跳台阶
// dp[i] = dp[i-1] + dp[i-2]
func jumpFloor(number int) int {
	m1, m2 := 1, 1
	for i := 0; i < number-1; i++ {
		m1 += m2
		m1, m2 = m2, m1
	}
	return m2
}
