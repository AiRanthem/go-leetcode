package top100

func majorityElement(nums []int) int {
	maj, time := 0, 0
	for i := 0; i < len(nums); i++ {
		if time == 0 {
			maj, time = nums[i], 0
		}
		if nums[i] == maj {
			time++
		} else {
			time--
		}
	}
	return maj
}
