package leetcode

//func maximalSquare(matrix [][]byte) int {
//	m, n := len(matrix), len(matrix[0])+1
//	mono := Stack{-1}
//	heights := make([]int, n)
//	left := make([]int, n)
//	result := 0
//	for i := 0; i < m; i++ {
//		matrix[i] = append(matrix[i], '0')
//		for j := 0; j < n; j++ {
//			if matrix[i][j] == '1' {
//				heights[j]++
//			} else {
//				heights[j] = 0
//			}
//			for len(mono) > 1 && heights[mono.Top()] >= heights[j] {
//				toCalcu := mono.Pop()
//				area := j - left[toCalcu] - 1
//				if area > heights[toCalcu] {
//					area = heights[toCalcu]
//				}
//				if result < area {
//					result = area
//				}
//			}
//			left[j] = mono.Top()
//			mono.Push(j)
//		}
//		mono.Pop()
//	}
//	return result * result
//}

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := new2dSlice(m+1, n+1)
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				dp[i+1][j+1] = 0
			} else {
				dp[i+1][j+1] = dp[i][j]
				if dp[i+1][j] < dp[i+1][j+1] {
					dp[i+1][j+1] = dp[i+1][j]
				}
				if dp[i][j+1] < dp[i+1][j+1] {
					dp[i+1][j+1] = dp[i][j+1]
				}
				dp[i+1][j+1]++
				if dp[i+1][j+1] > max {
					max = dp[i+1][j+1]
				}
			}
		}
	}
	return max * max
}
