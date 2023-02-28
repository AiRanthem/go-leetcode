package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		bytes := []byte(strs[i])
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		resultMap[string(bytes)] = append(resultMap[string(bytes)], strs[i])
	}
	var result [][]string
	for _, row := range resultMap {
		result = append(result, row)
	}
	return result
}
