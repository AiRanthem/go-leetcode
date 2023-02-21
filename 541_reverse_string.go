package leetcode

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	n := len(s) - 1
	for i := 0; i < len(s); i += k * 2 {
		if i+k > n {
			revers(bytes, i, n)
		} else {
			revers(bytes, i, i+k-1)
		}
	}
	return string(bytes)
}

func revers(s []byte, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
