package leetcode

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var canVisit func(x, y int) bool
	canVisit = func(x, y int) bool {
		return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '1'
	}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if !canVisit(x, y) {
			return
		}
		grid[x][y] = '2'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	var num int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				num++
				dfs(i, j)
			}
		}
	}
	return num
}
