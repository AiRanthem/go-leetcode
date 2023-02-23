package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	counter := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		counter[nums1[i]] = false
	}
	for i := 0; i < len(nums2); i++ {
		if _, ok := counter[nums2[i]]; ok {
			counter[nums2[i]] = true
		}
	}
	var result []int
	for i, t := range counter {
		if t {
			result = append(result, i)
		}
	}
	return result
}
