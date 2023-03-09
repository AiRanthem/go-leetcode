package leetcode

func findDuplicate(nums []int) int {
	for nums[0] != nums[nums[0]] {
		nums[0], nums[nums[0]] = nums[nums[0]], nums[0]
	}
	return nums[0]
}
