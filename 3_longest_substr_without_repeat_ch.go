package leetcode

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]struct{})
	i, result := 0, 0
	ss := []byte(s)
	for j, c := range ss {
		if _, ok := m[c]; ok {
			for ok {
				delete(m, ss[i])
				i++
				_, ok = m[c]
			}
			m[c] = struct{}{}
		} else {
			m[c] = struct{}{}
			result = maxInt(result, j-i+1)
		}
	}
	return result
}
