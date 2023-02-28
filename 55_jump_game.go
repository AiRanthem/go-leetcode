package leetcode

func canJump(nums []int) bool {
	i, j := 0, 0
	for {
		if i+nums[i] > j {
			j = i + nums[i]
		}
		if j >= len(nums)-1 {
			return true
		}
		i++
		if i > j {
			return false
		}
	}
}
