package NC61

// 两数之和
func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		m[numbers[i]] = i
	}
	for i := 0; i < len(numbers); i++ {
		if j, ok := m[target-numbers[i]]; ok && i != j {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}
