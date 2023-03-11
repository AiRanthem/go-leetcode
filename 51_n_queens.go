package leetcode

import "strings"

func solveNQueens(n int) (result [][]string) {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	var dfs func(int)
	dfs = func(row int) { // kth row
		if row == n {
			r := make([]string, n)
			for i := 0; i < n; i++ {
				r[i] = strings.Join(board[i], "")
			}
			result = append(result, r)
			return
		}
		for col := 0; col < n; col++ {
			ok := true
			for i := row - 1; ok && i >= 0; i-- {
				if board[i][col] == "Q" {
					ok = false
				}
			}
			for i, j := row-1, col-1; ok && i >= 0 && j >= 0; i, j = i-1, j-1 {
				if board[i][j] == "Q" {
					ok = false
				}
			}
			for i, j := row-1, col+1; ok && i >= 0 && j < n; i, j = i-1, j+1 {
				if board[i][j] == "Q" {
					ok = false
				}
			}
			if ok {
				board[row][col] = "Q"
				dfs(row + 1)
				board[row][col] = "."
			}
		}
	}
	dfs(0)
	return
}
