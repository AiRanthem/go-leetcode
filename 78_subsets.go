package leetcode

func subsets(nums []int) (ans [][]int) {
	var subset []int
	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			r := make([]int, len(subset))
			copy(r, subset)
			ans = append(ans, r)
			return
		}
		dfs(i + 1)                       // not put
		subset = append(subset, nums[i]) // put
		dfs(i + 1)
		subset = subset[:len(subset)-1]
	}
	dfs(0)
	return
}
