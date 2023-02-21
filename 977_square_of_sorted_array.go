package leetcode

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	i, j, k := 0, len(nums)-1, len(nums)-1
	for i <= j {
		l, r := nums[i]*nums[i], nums[j]*nums[j]
		if l < r {
			result[k] = r
			j--
		} else {
			result[k] = l
			i++
		}
		k--
	}
	return result
}
