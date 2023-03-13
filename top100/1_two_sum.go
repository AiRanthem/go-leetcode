package top100

func twoSum(nums []int, target int) (results []int) {
	set := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if j, ok := set[target-nums[i]]; ok && i != j {
			results = append(results, i)
			results = append(results, j)
			break
		}
	}
	return
}
