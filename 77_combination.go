package leetcode

func combine(n int, k int) [][]int {
	var results [][]int
	var path []int
	var dfs func(start int)
	dfs = func(start int) {
		path = append(path, start)
		if len(path) == k {
			r := make([]int, k)
			copy(r, path)
			results = append(results, r)
		} else {
			for i := start + 1; i <= n; i++ {
				dfs(i)
			}
		}
		path = path[:len(path)-1]
	}
	return results
}
