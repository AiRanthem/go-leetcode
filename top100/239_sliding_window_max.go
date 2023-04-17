package top100

func maxSlidingWindow(nums []int, k int) (result []int) {
	ms := MonoStack{}
	for i := 0; i < k-1; i++ {
		ms.RPush(nums[i])
	}
	for i := k - 1; i < len(nums); i++ {
		ms.RPush(nums[i])
		result = append(result, ms.Left())
		if ms.Left() == nums[i-k+1] {
			ms.LPop()
		}
	}
	return
}
