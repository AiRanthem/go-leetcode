package leetcode

func majorityElement(nums []int) int {
	master, times := 0, 0
	for i := 0; i < len(nums); i++ {
		if times == 0 {
			master = nums[i]
		}
		if nums[i] == master {
			times++
		} else {
			times--
		}
	}
	return master
}
