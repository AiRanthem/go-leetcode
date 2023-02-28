package leetcode

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}
	stack := make([]int, n*2)
	top := -1
	for i := 0; i < n*2; i++ {
		for top >= 0 && nums[stack[top]] < nums[i%n] {
			result[stack[top]] = nums[i%n]
			top--
		}
		top++
		stack[top] = i % n
	}
	return result
}
