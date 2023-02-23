package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		if tm := m[t[i]]; tm > 0 {
			m[t[i]] = tm - 1
		} else {
			return false
		}
	}
	return true
}
