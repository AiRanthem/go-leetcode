package leetcode

func commonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}
	resultMap := countChars(words[0])
	for i := 1; i < len(words); i++ {
		counter := countChars(words[i])
		for j := 0; j < 26; j++ {
			resultMap[j] = minInt(resultMap[j], counter[j])
		}
	}
	var result []string
	for i := 0; i < 26; i++ {
		for j := 0; j < resultMap[i]; j++ {
			result = append(result, string(rune(i+'a')))
		}
	}
	return result
}

func countChars(s string) []int {
	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
	}
	return counter
}
