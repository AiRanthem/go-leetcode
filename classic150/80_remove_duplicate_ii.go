package classic150

func removeDuplicates(nums []int) int {
	var i, j int
	for j < len(nums) {
		nums[i] = nums[j]
		if i > 0 && nums[i] == nums[i-1] {
			for j < len(nums) && nums[j] == nums[i] {
				j++
			}
		} else {
			j++
		}
		i++
	}
	return i
}
