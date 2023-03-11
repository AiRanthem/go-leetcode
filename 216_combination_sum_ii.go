package leetcode

func combinationSum3(k int, n int) (result [][]int) {
	var path []int
	var dfs func(start int, n int)
	dfs = func(start int, n int) {
		path = append(path, start)
		if len(path) == k {
			if n == 0 {
				r := make([]int, k)
				copy(r, path)
				result = append(result, r)
			}
		} else {
			for i := start + 1; i <= 9 && i <= n; i++ {
				dfs(i, n-i)
			}
		}
		path = path[:len(path)-1]
	}
	for i := 1; i <= 9; i++ {
		dfs(i, n-i)
	}
	return
}
