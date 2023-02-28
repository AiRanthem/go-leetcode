package leetcode

/*
84 56 35 20 10 4 1
28 21 15 10 6  3 1
7  6  5  4  3  2 1
1  1  1  1  1  1 1

C(down+right, down)
C(9, 3) = (9*8*7)/(3*2)
C(8,2) = (8*7)/(2) = 28
*/

//func uniquePaths(m int, n int) int {
//	if m > n {
//		return uniquePaths(n, m)
//	}
//	var big, small = int64(m + n - 2), int64(m - 1)
//	var up, down int64 = 1, 1
//	for i := 0; i < m-1; i++ {
//		up *= big
//		big--
//	}
//	for small > 1 {
//		down *= small
//		small--
//	}
//	return int(up / down)
//}

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
