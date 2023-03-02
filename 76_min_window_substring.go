package leetcode

func minWindow(s string, t string) string {
	want := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		want[t[i]]++
	}
	subLenLeft := len(t)
	l := 0
	result := s + "0"
	for r := 0; r < len(s); r++ {
		// always include right
		if w := want[s[r]]; w > 0 {
			subLenLeft--
		}
		want[s[r]]--
		// shrink left if full
		if subLenLeft == 0 {
			for want[s[l]] < 0 {
				want[s[l]]++
				l++
			}
			if r-l+1 < len(result) {
				result = s[l : r+1]
			}
			// release left bound
			want[s[l]]++
			subLenLeft++
			l++
		}
	}
	if len(result) > len(s) {
		return ""
	} else {
		return result
	}
}
