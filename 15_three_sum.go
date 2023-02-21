package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var result [][]int
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		target := -1 * nums[i]
		j, k := i+1, n-1
		for j < k {
			for j < k && j > i+1 && nums[j] == nums[j-1] {
				j++
			}
			for j < k && k < n-1 && nums[k] == nums[k+1] {
				k--
			}
			if j >= k {
				break
			}
			s := nums[j] + nums[k]
			if s == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
			} else if s > target {
				k--
			} else {
				j++
			}
		}
	}
	return result
}
