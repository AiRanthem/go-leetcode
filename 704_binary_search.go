package leetcode

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		p := (left + right) / 2
		if nums[p] == target {
			return p
		}
		if nums[p] > target {
			right = p - 1
		} else {
			left = p + 1
		}
	}
	return -1
}
