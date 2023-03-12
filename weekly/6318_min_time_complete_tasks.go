package weekly

import "sort"

func findMinimumTime(tasks [][]int) (cnt int) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	alloc := make([]int, 2001)
	for i := 0; i < len(tasks); i++ {
		start, end, duration := tasks[i][0], tasks[i][1], tasks[i][2]
		time := sum(alloc[start : end+1])
		for time < duration {
			time += alloc[end]
			alloc[end] += 1 - alloc[end]
			if alloc[end] == 0 {
				alloc[end] = 1
				time++
			}
			end--
		}
	}
	return sum(alloc)
}

func sum(arr []int) (s int) {
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return
}
