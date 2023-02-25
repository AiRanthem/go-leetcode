package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}
	l, r := 0, len(nums)-1
	result := []int{-1, -1}
	for l < r {
		m := (l + r) / 2
		if nums[m] >= target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	if l > len(nums)-2 {
		l = len(nums) - 2
	}
	if nums[l] == target {
		result[0] = l
	} else if nums[l+1] == target {
		result[0] = l + 1
	}
	l, r = 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] <= target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if r < 1 {
		r = 1
	}
	if nums[r] == target {
		result[1] = r
	} else if nums[r-1] == target {
		result[1] = r - 1
	}
	return result
}
