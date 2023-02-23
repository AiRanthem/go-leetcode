package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for one := 0; one < len(nums); one++ {
		if one > 0 && nums[one] == nums[one-1] {
			continue
		}
		for two := one + 1; two < len(nums); two++ {
			if two > one+1 && nums[two] == nums[two-1] {
				continue
			}
			three, four := two+1, len(nums)-1
			for three < four {
				if three > two+1 && nums[three] == nums[three]-1 {
					three++
					continue
				}
				if four < len(nums)-1 && nums[four] == nums[four+1] {
					four--
					continue
				}
				if s := nums[one] + nums[two] + nums[three] + nums[four]; s == target {
					result = append(result, []int{nums[one], nums[two], nums[three], nums[four]})
					four--
				} else if s > target {
					four--
				} else {
					three++
				}
			}
		}
	}
	return result
}
