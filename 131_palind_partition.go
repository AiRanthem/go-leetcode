package leetcode

func partition(s string) (result [][]string) {
	isPalind := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		row := make([]bool, len(s))
		row[i] = true
		if i+1 < len(s) && s[i] == s[i+1] {
			row[i+1] = true
		}
		isPalind[i] = row
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 2; j < len(s); j++ {
			isPalind[i][j] = isPalind[i+1][j-1] && s[i] == s[j]
		}
	}

	var part []string
	var dfs func(start int)
	dfs = func(start int) {
		if start >= len(s) {
			r := make([]string, len(part))
			copy(r, part)
			result = append(result, r)
		} else {
			for i := start; i < len(s); i++ {
				if isPalind[start][i] {
					part = append(part, s[start:i+1])
					dfs(i + 1)
					part = part[:len(part)-1]
				}
			}
		}
	}
	dfs(0)
	return
}
