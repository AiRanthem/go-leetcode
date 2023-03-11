package leetcode

import "sort"

func combinationSum2(candidates []int, target int) (result [][]int) {
	sort.Ints(candidates)
	var comb []int
	var dfs func(start, target int)
	dfs = func(start, target int) {
		if target == 0 {
			r := make([]int, len(comb))
			copy(r, comb)
			result = append(result, r)
		} else if target > 0 {
			for i := start; i < len(candidates); i++ {
				if i > start && candidates[i] == candidates[i-1] {
					continue
				}
				comb = append(comb, candidates[i])
				dfs(i+1, target-candidates[i])
				comb = comb[:len(comb)-1]
			}
		}
	}
	dfs(0, target)
	return
}
