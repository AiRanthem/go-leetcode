package classic150

func rotate(nums []int, k int) {
	k %= len(nums)
	var reverse func(i int, j int)
	reverse = func(i int, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	reverse(0, len(nums)-1)
	reverse(0, k-1)
	reverse(k, len(nums)-1)
}
