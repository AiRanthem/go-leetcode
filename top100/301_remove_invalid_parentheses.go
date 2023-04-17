package top100

import "math"

func removeInvalidParentheses(s string) []string {
	prefix := make([]int, len(s))
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			sum++
		} else if s[i] == ')' {
			sum--
		}
		prefix[i] = sum
	}

	results := make(map[string]struct{})
	var path []byte
	minDeletions := math.MaxInt
	var dfs func(i int, offset int, deletions int)
	dfs = func(i int, offset int, deletions int) {
		if deletions > minDeletions {
			return
		}
		if i == len(s) {
			if offset+prefix[i-1] != 0 {
				return
			}
			if deletions < minDeletions {
				minDeletions = deletions
				results = make(map[string]struct{})
			}
			r := string(path)
			if _, ok := results[r]; !ok {
				results[r] = struct{}{}
			}
			return
		}
		p := prefix[i] + offset
		switch s[i] {
		default:
			if p < 0 {
				return
			}
			path = append(path, s[i])
			dfs(i+1, offset, deletions)
			path = path[:len(path)-1]
		case '(':
			dfs(i+1, offset-1, deletions+1)
			path = append(path, '(')
			dfs(i+1, offset, deletions)
			path = path[:len(path)-1]
		case ')':
			dfs(i+1, offset+1, deletions+1)
			if p >= 0 {
				path = append(path, ')')
				dfs(i+1, offset, deletions)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0, 0, 0)
	var ans []string
	for k, _ := range results {
		ans = append(ans, k)
	}
	return ans
}

func removeInvalidParentheses2(s string) []string {
	prefix := make([]int, len(s))
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			sum++
		} else if s[i] == ')' {
			sum--
		}
		prefix[i] = sum
	}
	var dfs func(i int, offset int, deletions int)
	var buf []byte
	minDeletions := math.MaxInt
	results := make(map[string]struct{})
	dfs = func(i int, offset int, deletions int) {
		if deletions > minDeletions {
			return
		}
		if i == len(s) {
			if offset+prefix[i-1] != 0 {
				return
			}
			if deletions < minDeletions {
				minDeletions = deletions
				results = make(map[string]struct{})
			}
			results[string(buf)] = struct{}{}
			return
		}
		balance := offset + prefix[i]
		switch s[i] {
		case '(':
			dfs(i+1, offset-1, deletions+1)
			buf = append(buf, '(')
			dfs(i+1, offset, deletions)
			buf = buf[:len(buf)-1]
		case ')':
			dfs(i+1, offset+1, deletions+1)
			if balance >= 0 {
				buf = append(buf, ')')
				dfs(i+1, offset, deletions)
				buf = buf[:len(buf)-1]
			}
		default:
			if balance < 0 {
				return
			}
			buf = append(buf, s[i])
			dfs(i+1, offset, deletions)
			buf = buf[:len(buf)-1]
		}
	}
	dfs(0, 0, 0)
	var ans []string
	for r, _ := range results {
		ans = append(ans, r)
	}
	return ans
}
