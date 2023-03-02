package leetcode

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([]bool, m*n)
	var dfs func(int, int, int) bool
	dfs = func(x int, y int, process int) bool {
		if x >= 0 && x < m && y >= 0 && y < n && !visited[x*n+y] && board[x][y] == word[process] {
			visited[x*n+y] = true
			if process == len(word)-1 {
				return true
			}
			found := dfs(x+1, y, process+1) || dfs(x-1, y, process+1) || dfs(x, y+1, process+1) || dfs(x, y-1, process+1)
			visited[x*n+y] = false
			return found
		} else {
			return false
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
