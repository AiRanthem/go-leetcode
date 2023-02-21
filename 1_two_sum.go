package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		if pos, ok := m[target-num]; ok && i != pos {
			return []int{i, pos}
		}
	}
	return []int{}
}
