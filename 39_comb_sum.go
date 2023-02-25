package leetcode

import "sort"

var result39 [][]int

func combinationSum(candidates []int, target int) [][]int {
	result39 = [][]int{}
	sort.Ints(candidates)
	dfs39(candidates, 0, []int{}, target)
	return result39
}

func dfs39(candidates []int, start int, path []int, target int) {
	if target < 0 {
		return
	}
	if target == 0 {
		p := make([]int, len(path))
		copy(p, path)
		result39 = append(result39, p)
		return
	}
	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		dfs39(candidates, i, path, target-candidates[i])
		path = path[:len(path)-1]
	}
}
