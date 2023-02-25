package leetcode

//func search(nums []int, target int) int {
//	l, r := 0, len(nums)-1
//	for l <= r {
//		m := (l + r) / 2
//		if nums[m] == target {
//			return m
//		}
//		if nums[l] == target {
//			return l
//		}
//		if nums[r] == target {
//			return r
//		}
//		if nums[m] > nums[l] {
//			if nums[l] < target && target < nums[m] {
//				r = m - 1
//			} else {
//				l = m + 1
//			}
//		} else {
//			if nums[m] < target && target < nums[r] {
//				l = m + 1
//			} else {
//				r = m - 1
//			}
//		}
//	}
//	return -1
//}
