package leetcode

func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	p := nums[0]
	for l < r {
		for l < r && nums[r] < p {
			r--
		}
		for l < r && nums[l] >= p {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[l], nums[0] = nums[0], nums[l]
	if l == k-1 {
		return nums[l]
	} else if l > k-1 {
		return findKthLargest(nums[:l], k)
	} else {
		return findKthLargest(nums[l+1:], k-l-1)
	}
}
