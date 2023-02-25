package leetcode

var result46 [][]int

func permute(nums []int) [][]int {
	result46 = [][]int{}
	dfs46(nums, []int{}, map[int]struct{}{})
	return result46
}

func dfs46(nums []int, path []int, used map[int]struct{}) {
	n := len(path)
	if n == len(nums) {
		buf := make([]int, n)
		copy(buf, path)
		result46 = append(result46, buf)
		return
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if _, ok := used[num]; ok {
			continue
		}
		used[num] = struct{}{}
		path = append(path, num)
		dfs46(nums, path, used)
		path = path[:n]
		delete(used, num)
	}
}
