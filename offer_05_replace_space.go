package leetcode

func replaceSpace(st string) string {
	s := []byte(st)
	for i := 0; i < len(st); i++ {
		if s[i] == ' ' {
			s = append(s, 0, 0)
		}
	}
	i, j := len(st)-1, len(s)-1
	for i >= 0 {
		if s[i] != ' ' {
			s[j] = s[i]
			j--
			i--
		} else {
			s[j] = '0'
			s[j-1] = '2'
			s[j-2] = '%'
			j -= 3
			i--
		}
	}
	return string(s)
}
