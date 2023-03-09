package leetcode

/*
dp[i][j], 以第i个block结尾的j个黑色要涂几次
dp[i][j] = dp[i-1][j-1] + (block[i]=='W')
*/
// func minimumRecolors(blocks string, k int) int {
//	m := len(blocks)
//	dp := make([]int, m+1)
//	for i := 0; i < k; i++ {
//		for j := m; j > i; j-- {
//			dp[j] = dp[j-1]
//			if blocks[j-1] == 'W' {
//				dp[j]++
//			}
//		}
//	}
//	result := math.MaxInt
//	for i := k; i < m+1; i++ {
//		if result > dp[i] {
//			result = dp[i]
//		}
//	}
//	return result
// }

func minimumRecolors(blocks string, k int) int {
	l, r, cnt, min := 0, 0, 0, 0
	for r < k {
		if blocks[r] == 'W' {
			cnt++
		}
		r++
	}
	min = cnt
	for r < len(blocks) {
		if blocks[r] == 'W' {
			cnt++
		}
		if blocks[l] == 'W' {
			cnt--
		}
		r++
		l++
		if cnt < min {
			min = cnt
		}
	}
	return min
}
