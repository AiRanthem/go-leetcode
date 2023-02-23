package leetcode

func reverseWords(st string) string {
	s := []byte(st)
	// trim
	i, j := 0, len(s)-1
	for i <= j && (s[i] == ' ' || s[j] == ' ') {
		if s[i] == ' ' {
			i++
		}
		if s[j] == ' ' {
			j--
		}
	}
	s = s[i : j+1]
	// remove extra spaces
	i, j = 0, 0
	for j < len(s) {
		if s[j] == ' ' {
			for s[j+1] == ' ' {
				j++
			}
		}
		s[i] = s[j]
		i++
		j++
	}
	s = s[:i]
	// reverse
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	i = 0
	for i < len(s) {
		if s[i] == ' ' {
			i++
			continue
		}
		j = i + 1
		for j < len(s) && s[j] != ' ' {
			j++
		}
		next := j
		j--
		for i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
		i = next
	}
	return string(s)
}
