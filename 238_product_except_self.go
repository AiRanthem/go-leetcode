package leetcode

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	x := 1
	for i := 0; i < n; i++ {
		result[i] = x
		x *= nums[i]
	}
	x = 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= x
		x *= nums[i]
	}
	return result
}
