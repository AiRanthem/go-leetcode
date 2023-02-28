package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	var result [][]int
	l, r := intervals[len(intervals)-1][0], intervals[len(intervals)-1][1]
	for i := len(intervals) - 2; i >= 0; i-- {
		if intervals[i][1] < l {
			result = append(result, []int{l, r})
			l, r = intervals[i][0], intervals[i][1]
		}
		if intervals[i][0] < l {
			l = intervals[i][0]
		}
	}
	result = append(result, []int{l, r})
	return result
}
